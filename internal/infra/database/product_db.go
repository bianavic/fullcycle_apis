package database

import (
	"github.com/bianavic/fullcycle_apis/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

/*
page - o número da página que queremos exibir
limit - quantos registros serão exibidos por página
offset - especifica quantos registros devem ser ignorados antes de começar a exibir os resultados. usado para pular os registros das páginas anteriores.
pagina == 0 - primeira pagina
(page - 1) * limit - proximo offset = quantidade de paginas e multiplica - quais registros devem ser ignorados para exibir uma página específica de resultados
*/
func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}
	return products, err
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		// se erro produto nao existia, entao nao tem update
		return err
	}
	// se existe, atualiza
	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}
	// se existe, deleta
	return p.DB.Delete(product).Error
}
