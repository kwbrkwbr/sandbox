package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

type Member struct {
	Id   string // FirestoreのドキュメントID
	Name string
	Age  int32
}

// FetchAllMembers Memberコレクションに存在する全てのドキュメントを取得し、Member構造体のPスライスに入れて返す
func FetchAllMembers(ctx context.Context, client *firestore.Client) ([]*Member, error) {
	var members []*Member

	docRefs := client.Collection("Members").Documents(ctx)
	docs, err := docRefs.GetAll()
	if err != nil {
		return nil, err
	}
	for _, doc := range docs {
		var member *Member
		if err := doc.DataTo(&member); err != nil {
			return nil, err
		}
		member.Id = doc.Ref.ID
		members = append(members, member)
	}

	return members, nil
}

// PutMember 渡されたMember構造体の中身を、ドキュメントとして保存する
func PutMember(ctx context.Context, client *firestore.Client, member *Member) {
	_, _, err := client.Collection("Members").Add(ctx, map[string]interface{}{
		"name": member.Name,
		"age":  member.Age,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
}

// DeleteAllMember Memberコレクションに存在する全てのドキュメントを削除する
func DeleteAllMember(ctx context.Context, client *firestore.Client) {
	docRefs := client.Collection("Members").Documents(ctx)
	docs, err := docRefs.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, doc := range docs {
		doc.Ref.Delete(ctx)
	}
}

// DeleteMember ドキュメントIDを受け取って、単体のドキュメントを削除する
func DeleteMember(ctx context.Context, client *firestore.Client, id string) {
	docRef := client.Collection("Members").Doc(id)
	_, err := docRef.Delete(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
