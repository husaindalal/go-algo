package hash

import (
	"fmt"
	"strings"
)

type File struct {
	DirName  string
	FileName string
	Content  string
}

func FindDuplicate(files []File) [][]File {
	fileByContent := map[string][]File{}
	for _, f := range files {
		if fileByContent[f.Content] == nil {
			fileByContent[f.Content] = []File{}
		}
		fileByContent[f.Content] = append(fileByContent[f.Content], f)
	}

	result := [][]File{}
	for _, v := range fileByContent {
		result = append(result, v)
	}
	fmt.Println(result)
	return result
}

func FindDuplicateMain(paths []string) [][]string {
	files := []File{}
	for _, path := range paths {
		fileCount := strings.Count(path, " ")
		dirName := path[0:strings.Index(path, " ")]
		path = strings.TrimPrefix(path, dirName+" ")
		for i := 0; i < fileCount; i++ {
			idx := strings.Index(path, " ")
			if idx == -1 {
				idx = len(path)
			}
			fileNameContent := path[0:idx]

			txtIndex := strings.Index(fileNameContent, ".txt") + 4
			fileName := fileNameContent[0:txtIndex]
			content := fileNameContent[txtIndex+1 : len(fileNameContent)-1]
			files = append(files, File{
				DirName:  dirName,
				FileName: fileName,
				Content:  content,
			})

			path = strings.TrimPrefix(path, fileNameContent+" ")
		}
	}

	fmt.Println(files)

	return nil

}

/*

	result := hash.FindDuplicate([]hash.File{
		{
			DirName:  "a",
			FileName: "1.txt",
			Content:  "abc",
		},
		{
			DirName:  "a",
			FileName: "2.txt",
			Content:  "pqr",
		},
		{
			DirName:  "b",
			FileName: "3.txt",
			Content:  "pqr",
		},
		{
			DirName:  "b",
			FileName: "4.txt",
			Content:  "rfj",
		},
	})


	fmt.Printf("result %v %v %v %v\n", result, "x","y","z")

	r2 := hash.FindDuplicateMain([]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"})
	fmt.Printf("result %v %v %v %v\n", r2, "x","y","z")
*/
