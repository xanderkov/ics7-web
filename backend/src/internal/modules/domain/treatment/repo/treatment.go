package repo

import (
	"context"
	"hospital/internal/modules/db"
	"hospital/internal/modules/db/ent"
	"hospital/internal/modules/domain/treatment/dto"
)

type TreatmentRepo struct {
	client *ent.Client
}

func NewTreatmentRepo(client *ent.Client) *TreatmentRepo {
	return &TreatmentRepo{
		client: client,
	}
}

func (r *TreatmentRepo) GetById(ctx context.Context, id int) (*dto.Treatment, error) {
	Treatment, err := r.client.Treatment.Get(ctx, id)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToTreatmentDTO(Treatment), nil
}

func (r *TreatmentRepo) List(ctx context.Context) (dto.Treatments, error) {
	Treatments, err := r.client.Treatment.Query().All(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToTreatmentDTOs(Treatments), nil
}

func (r *TreatmentRepo) Create(ctx context.Context, dtm *dto.CreateTreatment) (*dto.Treatment, error) {
	Treatment, err := r.client.Treatment.Create().
		SetTablets(dtm.Tablets).
		SetSurvey(dtm.Survey).
		SetPsychologicalTreatment(dtm.PsychologicalTreatment).
		SetPatientNumber(dtm.PatientNumber).
		SetUpdatedAt(dtm.UpdateAt).
		Save(ctx)

	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToTreatmentDTO(Treatment), nil
}

func (r *TreatmentRepo) Update(ctx context.Context, id int, dtm *dto.UpdateTreatment) (*dto.Treatment, error) {
	Treatment, err := r.client.Treatment.UpdateOneID(id).
		SetTablets(dtm.Tablets).
		SetSurvey(dtm.Survey).
		SetPsychologicalTreatment(dtm.PsychologicalTreatment).
		SetPatientNumber(dtm.PatientNumber).
		SetUpdatedAt(dtm.UpdateAt).
		Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToTreatmentDTO(Treatment), nil
}

func (r *TreatmentRepo) Delete(ctx context.Context, id int) error {
	err := r.client.Treatment.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return db.WrapError(err)
	}

	return nil
}

func (r *TreatmentRepo) Restore(ctx context.Context, id int) (*dto.Treatment, error) {
	Treatment, err := r.client.Treatment.UpdateOneID(id).Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToTreatmentDTO(Treatment), nil
}

func ToTreatmentDTO(model *ent.Treatment) *dto.Treatment {
	if model == nil {
		return nil
	}
	return &dto.Treatment{
		Id:                     model.ID,
		Tablets:                model.Tablets,
		PsychologicalTreatment: model.PsychologicalTreatment,
		Survey:                 model.Survey,
		PatientNumber:          model.PatientNumber,
		UpdateAt:               model.UpdatedAt,
	}
}

func ToTreatmentDTOs(models ent.Treatments) dto.Treatments {
	if models == nil {
		return nil
	}
	dtms := make(dto.Treatments, len(models))
	for i := range models {
		dtms[i] = ToTreatmentDTO(models[i])
	}
	return dtms
}
