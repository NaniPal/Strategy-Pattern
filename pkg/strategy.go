package pkg

type Strategy interface {
	Route(startpoint int, endpoint int)
}
