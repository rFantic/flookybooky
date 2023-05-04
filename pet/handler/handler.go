package handler

import (
	"context"
	"flookybooky/pet/ent"
	pb "flookybooky/pet/proto"

	"github.com/jinzhu/copier"
)

type PetHandler struct {
	pb.UnimplementedPetServiceServer
	client ent.Client
}

func NewPetHandler(client ent.Client) (*PetHandler, error) {
	return &PetHandler{
		client: client,
	}, nil
}

func (h *PetHandler) CreatePets(ctx context.Context, pet *pb.Pet) (*pb.Pet, error) {
	err := h.client.Pet.Create().
		SetName(pet.Name).
		SetSpecies(pet.Species).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return pet, nil
}

func (h *PetHandler) GetPets(ctx context.Context, req *pb.GetPetsRequest) (*pb.GetPetsResponse, error) {
	pets, err := h.client.Pet.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	res := &pb.GetPetsResponse{}
	err = copier.Copy(&res.Pets, &pets)
	if err != nil {
		return nil, err
	}
	return res, nil
}
