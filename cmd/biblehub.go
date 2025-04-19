package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mniak/biblia/cmd/internal"
	"github.com/mniak/biblia/pkg/biblehub"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func IntelinearIterator(initialChapter biblehub.ChapterID) func(yield func(biblehub.InterlinearChapter, error) bool) {
	return func(yield func(biblehub.InterlinearChapter, error) bool) {
		lastChapter := initialChapter
		for {
			log.Printf("Downloading %s %d", lastChapter.Book, lastChapter.Chapter)
			interlinearChapter, err := biblehub.DefaultScraper.GetInterlinearChapter(lastChapter)
			if !yield(interlinearChapter, err) {
				break
			}
			if interlinearChapter.Next == nil {
				break
			}
			lastChapter = *interlinearChapter.Next
		}
	}
}

func BibleHub() *cobra.Command {
	cmd := &cobra.Command{
		Use: "biblehub",
	}

	var outputDir string
	cmdDownloadInterlinear := &cobra.Command{
		Use: "download-interlinear",
		Run: func(cmd *cobra.Command, args []string) {
			cobra.CheckErr(os.MkdirAll(outputDir, 0o755))

			initialChapter := lo.FromPtrOr(internal.AppState.LastInterlinearChapter, biblehub.ChapterID{
				Book:    "genesis",
				Chapter: 1,
			})

			for chapter, err := range IntelinearIterator(initialChapter) {
				cobra.CheckErr(err)

				bookdir := filepath.Join(outputDir, chapter.Book)
				cobra.CheckErr(os.MkdirAll(bookdir, 0o755))

				filename := filepath.Join(bookdir, fmt.Sprintf("%d.json", chapter.Chapter))
				f, err := os.Create(filename)
				cobra.CheckErr(err)
				defer f.Close()

				enc := json.NewEncoder(f)
				enc.SetIndent("", "  ")
				err = enc.Encode(chapter)
				cobra.CheckErr(err)

				internal.AppState.LastInterlinearChapter = &chapter.ChapterID
				internal.AppState.Save()
			}
		},
	}
	cmdDownloadInterlinear.Flags().StringVarP(&outputDir, "output-dir", "o", "", "Output directory")
	cmdDownloadInterlinear.MarkFlagRequired("output-dir")

	cmd.AddCommand(cmdDownloadInterlinear)
	return cmd
}
