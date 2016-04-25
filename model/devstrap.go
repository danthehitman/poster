package model

type DevStrap struct {

}

func (ds DevStrap) InitDevDbRecords() {
	var err error

	user1 := User{FirstName:"Dan",
		LastName:"Frank",
		Email:"dan.frank@comcast.net",
		Password:"test",
		IsSuperUser:true,
	}
	user1, err = CreateUser(user1)
	checkErr(err, "Failed to create User1")

	user2 := User{FirstName:"Testy",
		LastName:"Testerton",
		Email:"test@test.com",
		Password:"test",
	}
	user2, err = CreateUser(user2)
	checkErr(err, "Failed to create User2")

	post1 := Post{
		Title:"Post1",
		Description:"Post1 Description",
		Body:"Post1 body.",
		Owner:user2,
	}
	post1, err = CreatePost(post1)
	checkErr(err, "Failed to create Post1")

	post2 := Post{
		Title:"Post2",
		Description:"Post2 Description",
		Body:"Post2 body.",
		Owner:user1,
	}
	post2, err = CreatePost(post2)
	checkErr(err, "Failed to create Post2")

	post3 := Post{
		Title:"Post3",
		Description:"Post3 Description",
		Body:"Post3 body.",
		Owner:user1,
	}
	post3, err = CreatePost(post3)
	checkErr(err, "Failed to create Post3")

	auth1 := ResourceAuthorization{
		UserId: user2.Uuid,
		ResourceId:post1.Uuid,
		Action:WriteResourceAction,
		ResourceType:"post",
	}
	err = CreateResourceAuthorization(auth1)
	checkErr(err, "Failed to create Auth1")

	auth2 := ResourceAuthorization{
		UserId: user1.Uuid,
		ResourceId:post2.Uuid,
		Action:WriteResourceAction,
		ResourceType:"post",
	}
	err = CreateResourceAuthorization(auth2)
	checkErr(err, "Failed to create Auth2")

	auth3 := ResourceAuthorization{
		UserId: user1.Uuid,
		ResourceId:post3.Uuid,
		Action:WriteResourceAction,
		ResourceType:"post",
	}
	err = CreateResourceAuthorization(auth3)
	checkErr(err, "Failed to create Auth3")

	authGroup1 := ResourceGroup{
		ParentResourceId: post1.Uuid,
		ResourceId: post2.Uuid,
		ResourceType:"post",
	}
	err = CreateResourceGroup(authGroup1)
	checkErr(err, "Failed to create AuthGroup1")
}