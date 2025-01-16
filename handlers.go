package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	log.Println(" test2")
	v := viewData{}
	// v.Id = func(s string) string {
	// 	return strings.ReplaceAll(s, " ", "_")
	// }

	v.Tags = append(v.Tags, tags...)
	v.Categories = append(v.Categories, cats...)
	exeTmpl(w, r, &v, "main.tmpl")
}

// submitRoot verifies a users submissions and then adds it to the database.
func submitData(w http.ResponseWriter, r *http.Request) {
	// 	mr, err := r.MultipartReader()
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	var bodyText string
	// 	var tempFileName string
	// 	var mediaType string
	// 	for {
	// 		part, err_part := mr.NextPart()
	// 		if err_part == io.EOF {
	// 			break
	// 		}
	// 		if part.FormName() == "myFile" {
	// 			fileBytes, err := io.ReadAll(io.LimitReader(part, 10<<20))
	// 			if err != nil {
	// 				log.Println(err)
	// 			}
	// 			mt := http.DetectContentType(fileBytes)
	// 			if mt != "application/pdf" {
	// 				ajaxResponse(w, map[string]string{
	// 					"success": "false",
	// 					"replyID": "",
	// 					"error":   "pdf only",
	// 				})
	// 				return
	// 			}
	// 			tempFile, err := os.CreateTemp("public/media/pdf", "u-*.pdf")
	// 			if err != nil {
	// 				log.Println(err)
	// 			}
	// 			defer tempFile.Close()
	// 			tempFileName = tempFile.Name()
	// 			tempFile.Write(fileBytes)
	// 		}
	// 		if part.FormName() == "myText" {
	// 			buf := new(bytes.Buffer)
	// 			buf.ReadFrom(part)
	// 			bodyText = buf.String()
	// 			log.Println("bt:", bodyText)
	// 		}
	// 		if part.FormName() == "parent" {
	// 			buf := new(bytes.Buffer)
	// 			buf.ReadFrom(part)
	// 		}
	// 	}

	//	var data *post = &post{
	//		Id:        genPostID(10),
	//		TS:        time.Now(),
	//		BodyText:  bodyText,
	//		MediaType: mediaType,
	//	}
	//
	// data.FTS = data.TS.Format("2006-01-02 03:04:05 pm")
	// // rdb.HSet(
	// // 	rdx, data.Id,
	// // 	"bodytext", bodyText,
	// // 	"id", data.Id,
	// // 	"ts", data.TS,
	// // 	"fts", data.FTS,
	// // 	"parent", parent,
	// // 	"childCount", "0",
	// // 	"media", tempFileName,
	// // 	"mediaType", mediaType,
	// // )
	// // rdb.ZAdd(rdx, "ANON:POSTS:CHRON", redis.Z{Score: float64(time.Now().UnixMilli()), Member: data.Id})
	// // rdb.ZAdd(rdx, "ANON:POSTS:RANK", redis.Z{Score: 0, Member: data.Id})
	// // popLast()
	//
	//	ajaxResponse(w, map[string]string{
	//		"success":   "true",
	//		"replyID":   data.Id,
	//		"timestamp": data.FTS,
	//		"tfn":       tempFileName,
	//	})
	//
	// // beginCache()
}
