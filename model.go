package main

import (
	"database/sql"
	"errors"
)

type Resource struct {
	ResId int `json:"ResId"`
	ResName string `json:"ResName"`
	ResDesc string `json:"ResDesc"`
}

func (p *Resource) getProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Resource) updateProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Resource) deleteProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Resource) createProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}


func getProducts(db *sql.DB, start, count int) ([]Resource, error) {
	return nil, errors.New("Not implemented")
}