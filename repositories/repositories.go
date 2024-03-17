package repositories

import (
	"company/models"
	"database/sql"
	"errors"
	"log"
)

type IRepositorie interface {
	AddRegisterRepositorie(register models.RequestRegister) error
	GetAllEmployeesRepositorie() ([]models.ResponseEmployee, error)
	GetCompanyByIDRepositorie(cpnID string) (models.ResponseCompany, error)
	UpdateEmployeeRepositorie(employee models.RequestUpdateEmployee) error
	DeleteEmployeeRepositorie(epyID string) error
	
}

type repositorie struct {
	db *sql.DB
}

func NewRepositorie(db *sql.DB) IRepositorie {
	return &repositorie{db: db}
}

func (r *repositorie) AddRegisterRepositorie(register models.RequestRegister) error {
	var cpnID string
	// ค้นหาบริษัทจากชื่อ
	err := r.db.QueryRow("SELECT cpn_id FROM company WHERE cpn_name = $1", register.CpnName).Scan(&cpnID)
	if err != nil {
		// หากไม่พบบริษัท ให้สร้างบริษัทใหม่
		if err == sql.ErrNoRows {
			err := r.db.QueryRow("INSERT INTO company (cpn_name, cpn_address) VALUES ($1, $2) RETURNING cpn_id", register.CpnName, register.CpnAddress).Scan(&cpnID)
			if err != nil {
				log.Printf("failed to insert into company. Error: %v", err)
				return errors.New("failed to insert into company table (API addRegister)")
			}
		} else {
			return errors.New("failed to query company table (API addRegister)")
		}
	}

	// ตรวจสอบว่ามีพนักงานชื่อซ้ำกันหรือไม่
	var empName string
	err = r.db.QueryRow("SELECT epy_name FROM employee WHERE cpn_id = $1 AND epy_name = $2", cpnID, register.EmpName).Scan(&empName)
	if err == nil {
		// หากมีชื่อซ้ำกัน ให้คืน error
		return errors.New("employee already exists")
	} else if err != sql.ErrNoRows {
		log.Printf("failed to query employee table. Error: %v", err)
		return errors.New("failed to query employee table (API addRegister)")
	}

	// เพิ่มข้อมูลพนักงานใหม่
	_, err = r.db.Exec("INSERT INTO employee (epy_name, epy_position, epy_email, epy_phone, cpn_id) VALUES ($1, $2, $3, $4, $5)", register.EmpName, register.EmpPosition, register.EmpEmail, register.EmpPhone, cpnID)

	if err != nil {
		log.Printf("failed to insert into employee table. Error: %v", err)
		return errors.New("failed to insert into employee table (API addRegister)")
	}

	return nil
}

func (r *repositorie) GetAllEmployeesRepositorie() ([]models.ResponseEmployee, error) {
	rows, err := r.db.Query("SELECT * FROM employee")
	if err != nil {
		log.Printf("failed to query employee table. Error: %v", err)
		return nil, errors.New("failed to query employee table (API GetAllEmployeesRepositorie)")
	}
	defer rows.Close()

	employees := []models.ResponseEmployee{}
	for rows.Next() {
		var employee models.ResponseEmployee
		if err := rows.Scan(&employee.EpyID, &employee.EpyName, &employee.EpyPosition, &employee.EpyEmail, &employee.EpyPhone, &employee.CpnID); err != nil {
			log.Printf("failed to scan row. Error: %v", err)
			return nil, errors.New("failed to scan row (API GetAllEmployeesRepositorie)")
		}
		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		log.Printf("error in iterating rows. Error: %v", err)
		return nil, errors.New("error in iterating rows (API GetAllEmployeesRepositorie)")
	}

	return employees, nil
}

func (r *repositorie) GetCompanyByIDRepositorie(cpnID string) (models.ResponseCompany, error) {
	var company models.ResponseCompany

	err := r.db.QueryRow("SELECT  cpn_name, cpn_address FROM company WHERE cpn_id = $1", cpnID).Scan(&company.CpnName, &company.CpnAddress)
	if err != nil {
		log.Printf("failed to query company by ID. Error: %v", err)
		return models.ResponseCompany{}, err
	}

	return company, nil
}
func (r *repositorie) UpdateEmployeeRepositorie(employee models.RequestUpdateEmployee) error {
	// var count int
	// err := r.db.QueryRow("SELECT COUNT(*) FROM employee WHERE epy_id = $1", employee.EpyID).Scan(&count)
	// if err != nil {
	// 	log.Printf("failed to query employee with ID %s. Error: %v", employee.EpyID, err)
	// 	return errors.New("failed to query employee with ID")
	// }
	// if count == 0 {
	// 	log.Printf("employee with ID %s not found", employee.EpyID)
	// 	return errors.New("employee not found")
	// }
	err := r.db.QueryRow("SELECT epy_id FROM employee WHERE epy_id = $1", employee.EpyID).Scan(&employee.EpyID)
	if err != nil {
		log.Printf("failed to find employee with ID %s. Error: %v", employee.EpyID, err)
		return errors.New("failed to find employee with ID")
	}

	// ทำการอัปเดตข้อมูลของ employee ที่พบ
	query := `
		UPDATE employee
		SET
			epy_name = COALESCE(NULLIF($2, ''), epy_name),
			epy_position = COALESCE(NULLIF($3, ''), epy_position),
			epy_email = COALESCE(NULLIF($4, ''), epy_email),
			epy_phone = COALESCE(NULLIF($5, ''), epy_phone)
		WHERE epy_id = $1
	`
	_, err = r.db.Exec(query, employee.EpyID, employee.EpyName, employee.EpyPosition, employee.EpyEmail, employee.EpyPhone)
	if err != nil {
		log.Printf("failed to update employee partially. Error: %v", err)
		return errors.New("failed to update employee partially")
	}
	return nil
}

func (r *repositorie) DeleteEmployeeRepositorie(epyID string) error {
	var foundEpyID string
	err := r.db.QueryRow("SELECT epy_id FROM employee WHERE epy_id = $1", epyID).Scan(&foundEpyID)
	if err != nil {
		log.Printf("failed to find employee with ID %s. Error: %v", epyID, err)
		return errors.New("failed to find employee with ID")
	}
  
    _, err = r.db.Exec("DELETE FROM employee WHERE epy_id = $1", foundEpyID)
    if err != nil {
        log.Printf("failed to delete employee with ID %s. Error: %v", epyID, err)
        return errors.New("failed to delete employee")
    }
    return nil
}


