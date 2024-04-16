package services

import (
	"bytes"
	"context"
	"errors"
	"finance/models"
	"finance/repository"
	"fmt"
	"time"

	// "io/ioutil"
	"os"
	"strconv"

	v7uuid "github.com/gofrs/uuid"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

type TransactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (ts *TransactionService) GetTransactionByID(ctx context.Context, id uuid.UUID) (*models.Transaction, error) {
	return ts.repo.GetTransactionByID(ctx, id)
}

func (ts *TransactionService) GetTransactions(ctx context.Context) ([]models.Transaction, error) {

	result, err := ts.repo.GetTransactions(ctx)
	if err != nil {
		return nil, err

	}

	return result, err
}

// func (ts *TransactionService) CreateTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {

// 	err := ts.repo.CreateTransaction(ctx, transaction)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return transaction, nil
// }

func (ts *TransactionService) CreateTransaction(ctx context.Context, filePath string) error {
	// Read Excel file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	fmt.Println("service:", filePath)

	// Open Excel file for reading using excelize

	r := bytes.NewReader(data)

	f, err := excelize.OpenReader(r) // Use the reader here

	if err != nil {
		return fmt.Errorf("error opening Excel file: %w", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("Error closing Excel file:", err)
		}
	}()

	// Get sheet name (adjust sheet name as needed)
	sheetName := f.GetSheetName(0)

	// Read user data from sheet
	var transactions []models.Transaction
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return fmt.Errorf("error getting rows from sheet: %w", err)
	}

	if len(rows) > 1 {
		for i := 1; i < len(rows); i++ {

			// layout := "01/2/2006" // Example layout for MM/DD/YYYY
			layout := "01-02-06" // Example layout for MM/DD/YYYY

			// transactionDate, err := time.Parse(layout, "02/17/2023")

			transactionDate, err := time.Parse(layout, rows[i][0])

			if err != nil {
				fmt.Println("Error parsing string:", err)
			}

			// payeeId, err := v7uuid.FromString(rows[i][3])
			// if err != nil {
			// 	fmt.Println("Error parsing string:", err)
			// }

			amount, err := strconv.ParseFloat(rows[i][8], 64) // Convert deposit to float64
			if err != nil {
				// Handle conversion error (e.g., log error, skip row)
				fmt.Printf("Error converting deposit (%s) to float64 in row %d: %w\n", rows[i][8], i+1, err)
				continue // Skip to the next row
			}

			clrStr := rows[i][7] // Get the cell value as a string
			var clr = false
			if clrStr == "" {
				clr = false
			} else {
				clr, err = strconv.ParseBool(clrStr) // Convert deposit to float64
				if err != nil {
					// Handle conversion error (e.g., log error, skip row)
					fmt.Printf("Error converting clr (%s) to float64 in row %d: %w\n", rows[i][7], i+1, err)
					continue // Skip to the next row
				}
			}

			registerId, err_reg_parse := v7uuid.FromString("00000000-0000-0000-0000-000000000001")
			if err_reg_parse != nil {
				fmt.Println("Error parsing string:", err)
			}

			account := rows[i][1]

			register, reg_err := ts.repo.GetRegisterByName(ctx, account)

			if reg_err != nil {

			}
			if register != nil {
				registerId = register.ID

			} else {
				register := models.Register{

					Name:   account,
					OrgID:  registerId,
					UserID: registerId,
				}

				ts.repo.CreateTransactionRegister(ctx, &register)

				register2, reg_err2 := ts.repo.GetRegisterByName(ctx, account)

				if reg_err2 != nil {

				}

				registerId = register2.ID

			}

			

			payeeId, err_payee_parse := v7uuid.FromString("00000000-0000-0000-0000-000000000001")
			if err_payee_parse != nil {
				fmt.Println("Error parsing string:", err)
			}

			description := rows[i][3]

			payee, payee_err := ts.repo.GetPayeeByName(ctx, description)

			if payee_err != nil {

			}

			if payee != nil {
				payeeId = payee.ID

			} else {

				payee_model := models.Payee{

					Name:   description,
					OrgID:  payeeId,
					UserID: payeeId,
				}

				a := ts.repo.CreateTransactionPayee(ctx, &payee_model)

				fmt.Println("m,", a)
				payee2, payee_err2 := ts.repo.GetPayeeByName(ctx, description)

				if payee_err2 != nil {

				}

				payeeId = payee2.ID

			}

			categoryId, err_cat_parse := v7uuid.FromString("00000000-0000-0000-0000-000000000001")
			if err_cat_parse != nil {
				fmt.Println("Error parsing string:", err)
			}

			category_cell := rows[i][5]

			category, category_err := ts.repo.GetTransactionCategoryByName(ctx, category_cell)

			if category_err != nil {

			}

			if category != nil {
				categoryId = category.ID

			} else {
				category_model := models.TransactionCategory{

					Name:           category_cell,
					CategoryTypeID: categoryId,
					Icon:           999999999999999,
					Color:          "Yellow",
					Hidden:         false,
					Comment:        "Good",
					OrgID:          categoryId,
					UserID:         categoryId,
				}

				ts.repo.CreateTransactionCategory(ctx, &category_model)

				category2, category_err2 := ts.repo.GetTransactionCategoryByName(ctx, category_cell)

				if category_err2 != nil {

				}

				categoryId = category2.ID

			}

			tagId, err_tag_parse := v7uuid.FromString("00000000-0000-0000-0000-000000000001")
			if err_tag_parse != nil {
				fmt.Println("Error parsing string:", err)
			}

			tag_cell_str := rows[i][6]

			var tag_cell = ""

			if tag_cell_str == "" {
				tag_cell = "Default"
			} else {
				tag_cell = tag_cell_str
			}

			tag, tag_err := ts.repo.GetTagByName(ctx, tag_cell)

			if tag_err != nil {

			}
			if tag != nil {
				tagId = tag.ID

			} else {
				tag_model := models.Tag{

					Name:   tag_cell,
					OrgID:  tagId,
					UserID: tagId,
				}

				ts.repo.CreateTransactionTag(ctx, &tag_model)

				tag2, tag_err2 := ts.repo.GetTagByName(ctx, tag_cell)

				fmt.Println("tag,", tag2)

				if tag_err2 != nil {

				}

				tagId = tag2.ID

			}

			if amount > 0 {
				transaction := models.Transaction{

					TransactionDate:       transactionDate,
					RegisterID:            registerId,
					ReferenceNumber:       rows[i][2],
					PayeeID:               payeeId,
					Memo:                  rows[i][4],
					TransactionCategoryID: categoryId,
					TagID:                 tagId,
					Clr:                   clr,
					Deposit:               amount,
				}
				transactions = append(transactions, transaction)
			} else {
				transaction := models.Transaction{

					TransactionDate:       transactionDate,
					RegisterID:            registerId,
					ReferenceNumber:       rows[i][2],
					PayeeID:               payeeId,
					Memo:                  rows[i][4],
					TransactionCategoryID: categoryId,
					TagID:                 tagId,
					Clr:                   clr,
					Payment:               amount,
				}
				transactions = append(transactions, transaction)
			}

		}
	} else {
		return errors.New("no transaction data found in Excel file")
	}

	// Store users in database (use the UserRepository)
	err = ts.repo.CreateTransaction(ctx, transactions)
	if err != nil {
		return fmt.Errorf("error creating transactions: %w", err)
	}

	return nil
}
func (ts *TransactionService) UpdateTransaction(ctx context.Context, id uuid.UUID, transaction *models.Transaction) (*models.Transaction, error) {
	err := ts.repo.UpdateTransaction(ctx, id, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil

}

func (ts *TransactionService) DeleteTransaction(ctx context.Context, id uuid.UUID) error {
	return ts.repo.DeleteTransaction(ctx, id)
}
