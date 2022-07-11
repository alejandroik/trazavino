package entity

type Bottling struct {
	process   Process
	cellar    Cellar
	bottleQty int
	wine      Wine
}
