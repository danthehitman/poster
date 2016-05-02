package model

import "time"

type DevStrap struct {

}

var (
	user1Guid string = "33952134-8b8b-4b27-98e0-875ad65bb0a0"
	user2Guid string = "ed4249fb-c29b-4bf8-b14e-93791047a20d"
	journal1Guid string = "994b4b2c-e4bb-4d6b-93cc-30015931c655"
	journal2Guid string = "2a7efb9a-ea8f-48be-aa1d-8def9684bf41"
	post1Guid string = "064dd22c-3724-4ee0-b07e-fcdc9ecad453"
	post2Guid string = "03764a77-d6bb-4751-b852-3c1d66af0591"
	post3Guid string = "f38de052-1c65-41be-abae-c55a409b63f2"
	image1Guid string = "dde0aaac-114b-4690-863e-d6435ae45b65"
	image2Guid string = "e63979aa-b10d-420d-8212-2dfab32f6dc1"
	image3Guid string = "725b7b6e-b6b2-4c65-b3a8-48b393dc9071"
	link1Guid string = "95dc1ec8-9b77-4334-a04d-e1eba4bb19e8"
	link2Guid string = "74a888ab-eccd-4879-9798-491f56641dd2"
	link3Guid string = "8775a758-aa33-4669-94b2-0895562e5193"
	session1Guid string = "9b62ae00-374e-4be8-b1ae-b5df75f388ea"

	user1 User
	user2 User
	journal1 Journal
	journal2 Journal
	post1 Post
	post2 Post
	post3 Post
	image1 Image
	image2 Image
	image3 Image
	link1 Link
	link2 Link
	link3 Link
)

func (ds DevStrap) InitDevDbRecords() {
	ds.createUsers()
	ds.createImages()
	ds.createLinks()
	ds.createPosts()
	ds.createJournals()
	ds.createResourceAuthorizationsAndGroups()
}

func (ds DevStrap) createUsers() {
	var err error

	user1 = User{FirstName:"Dan",
		LastName:"Frank",
		Email:"dan.frank@comcast.net",
		Password:"test",
		IsSuperUser:true,
		Uuid:user1Guid,
	}
	user1, err = CreateUser(user1)
	checkErr(err, "Failed to create User1")

	user2 = User{
		FirstName:"Testy",
		LastName:"Testerton",
		Email:"test@test.com",
		Password:"test",
		Uuid:user2Guid,
	}
	user2, err = CreateUser(user2)
	checkErr(err, "Failed to create User2")

	CreateSession(Session{Uuid:session1Guid, ExpirationDate:time.Now().Add(time.Hour * 24), UserId: user1Guid})
}

func (ds DevStrap) createJournals() {
	var err error

	journal1 = Journal{
		Uuid:journal1Guid,
		Title:"Journal 1",
		IsPublic:true,
		Description:"Journal 1 description.",
		OwnerId:user1Guid,
		Posts:[]Post{post1, post2},
		Images:[]Image{image1, image2},
		Links:[]Link{link1, link2},
	}
	journal1, err = CreateJournal(journal1)
	checkErr(err, "Failed to create Journal1")

	journal2 = Journal{
		Uuid:journal2Guid,
		Title:"Journal 2",
		IsPublic:false,
		Description:"Journal 2 description.",
		OwnerId:user1Guid,
		Posts:[]Post{post2, post3},
		Images:[]Image{image3},
		Links:[]Link{link3},
	}
	journal2, err = CreateJournal(journal2)
	checkErr(err, "Failed to create Journal2")
}

func (ds DevStrap) createPosts() {
	var err error

	post1 = Post{
		Uuid:post1Guid,
		Title:"Post1",
		Description:"Post1 Description",
		Body:"Post1 body.",
		OwnerId:user1Guid,
		Images:[]Image{image1, image2},
		Links:[]Link{link1, link2},
	}
	post1, err = CreatePost(post1)
	checkErr(err, "Failed to create Post1")

	post2 = Post{
		Uuid:post2Guid,
		Title:"Post2",
		Description:"Post2 Description",
		Body:"Post2 body.",
		OwnerId:user1Guid,
		Images:[]Image{image3},
		Links:[]Link{link3},
	}
	post2, err = CreatePost(post2)
	checkErr(err, "Failed to create Post2")

	post3 = Post{
		Uuid:post3Guid,
		Title:"Post3",
		Description:"Post3 Description",
		Body:"Post3 body.",
		OwnerId:user1Guid,
	}
	post3, err = CreatePost(post3)
	checkErr(err, "Failed to create Post3")
}

func (ds DevStrap) createImages() {
	var err error

	image1 = Image{
		Uuid:image1Guid,
		Title:"Image 1",
		Description:"Image 1 description.",
		File:"Image1.png",
		OwnerId:user1Guid,
	}
	image1, err = CreateImage(image1)
	checkErr(err, "Failed to create Image1")

	image2 = Image{
		Uuid:image2Guid,
		Title:"Image 2",
		Description:"Image 2 description.",
		File:"Image2.png",
		OwnerId:user1Guid,
	}
	image2, err = CreateImage(image2)
	checkErr(err, "Failed to create Image2")

	image3 = Image{
		Uuid:image3Guid,
		Title:"Image 3",
		Description:"Image 3 description.",
		File:"Image3.png",
		OwnerId:user1Guid,
	}
	image3, err = CreateImage(image3)
	checkErr(err, "Failed to create Image3")
}

func (ds DevStrap) createLinks() {
	var err error

	link1 = Link{
		Uuid:link1Guid,
		Title:"Link1",
		Description:"Link1 description.",
		Url:"http://google.com",
		OwnerId:user1Guid,
	}
	link1, err = CreateLink(link1)
	checkErr(err, "Failed to create Link1")

	link2 = Link{
		Uuid:link2Guid,
		Title:"Link2",
		Description:"Link2 description.",
		Url:"http://google.com",
		OwnerId:user1Guid,
	}
	link2, err = CreateLink(link2)
	checkErr(err, "Failed to create Link2")

	link3 = Link{
		Uuid:link3Guid,
		Title:"Link3",
		Description:"Link3 description.",
		Url:"http://google.com",
		OwnerId:user1Guid,
	}
	link3, err = CreateLink(link3)
	checkErr(err, "Failed to create Link3")
}

func (ds DevStrap) createResourceAuthorizationsAndGroups() {
	var err error

	auth1 := ResourceAuthorization{
		UserId: user2Guid,
		ResourceId:post1Guid,
		Action:EditResourceAction,
		ResourceType:"post",
	}
	err = CreateResourceAuthorization(auth1)
	checkErr(err, "Failed to create Auth1")

	auth2 := ResourceAuthorization{
		UserId: user1Guid,
		ResourceId:post2Guid,
		Action:EditResourceAction,
		ResourceType:"post",
	}
	err = CreateResourceAuthorization(auth2)
	checkErr(err, "Failed to create Auth2")

	auth3 := ResourceAuthorization{
		UserId: user1Guid,
		ResourceId:post3Guid,
		Action:EditResourceAction,
		ResourceType:"post",
	}
	err = CreateResourceAuthorization(auth3)
	checkErr(err, "Failed to create Auth3")

	authGroup1 := ResourceGroup{
		ParentResourceId: post1Guid,
		ResourceId: post2Guid,
		ResourceType:"post",
	}
	err = CreateResourceGroup(authGroup1)
	checkErr(err, "Failed to create AuthGroup1")
}

