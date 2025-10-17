

func HistoryCheckSummary(filters dto.CheckDetailsFilters) (*dto.CheckDetailsResult, error) {
	query := `
		SELECT 
			(SELECT COUNT(*) FROM check_products p WHERE cs.id = p.check_summary_id) AS check_products,
			cs.check_date, u.code, cp.product_sku, p.descr_product, p.last_date_update, cp.collected_price, cp.pdv,
			(
				(SELECT COUNT(*) FROM check_products p WHERE cs.id = p.check_summary_id AND (p.status_price = 'N/SKU' OR p.status_price = 'NOK' OR p.status_price = 'CK'))
				+
				(SELECT COUNT(*) FROM check_products p WHERE cs.id = p.check_summary_id AND (p.status_product = 'N/SKU' OR p.status_product = 'NOK' OR p.status_product = 'CK'))
			)  AS occurrence,
			CASE
				WHEN cp.status_product = 'OK' THEN 'CHECKED'
				WHEN cp.status_product = 'CK' THEN 'UNABLE_READ'
				WHEN cp.status_product = 'NOK' THEN 'OUT_OF_POSITION'
				WHEN cp.status_product = 'N/SKU' THEN 'NOT_FOUND'
			END AS status_product,
				CASE
				WHEN cp.status_price = 'OK' THEN 'CHECKED'
				WHEN cp.status_price = 'CK' THEN 'UNABLE_READ'
				WHEN cp.status_price = 'NOK' THEN 'WRONG_PRICE'
				WHEN cp.status_price = 'N/SKU' THEN 'SKU_NOT_FOUND'
			END AS status_label,
			pe.ean
			FROM check_summary cs INNER JOIN check_products cp ON cs.id=cp.check_summary_id
			INNER JOIN users u ON cs.user_id =u.id
			INNER JOIN products p ON cp.product_sku = p.sku
			INNER JOIN product_ean pe ON p.sku = pe.product_sku
			WHERE cs.id = $1 AND pe.is_primary = 1
       		
	`
	var image []byte
	queryImage := `
		SELECT image FROM check_images WHERE check_summary_id = $1
	`

	_ = DB.QueryRow(queryImage, filters.CheckSummaryID).Scan(&image)

	argIdx := 2
	var args []interface{}

	args = append(args, filters.CheckSummaryID)
	var condition = []string{"", "", "AND (", "OR"}

	if filters.ProductStatus != "" {
		query += fmt.Sprintf(" %s cp.status_product= $%d", condition[argIdx], argIdx)
		args = append(args, filters.ProductStatus)
		argIdx++
	}

	if filters.TicketStatus != "" {
		query += fmt.Sprintf(" %s cp.status_price= $%d", condition[argIdx], argIdx)
		args = append(args, filters.TicketStatus)
		argIdx++
	}

	if argIdx > 2 {
		query += " ) "
	}

	query += " ORDER BY cp.product_sku ASC"

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result *dto.CheckDetailsResult
	var items []*dto.CheckDetailsItem
	var totalCheckedItems, totalOccurrences int
	var checkDate time.Time
	var effectiveDate time.Time
	var workerCode int
	var productSKU, description string
	var checkedPrice, pdv float64
	var productStatus, ticketStatus, ean string

	for rows.Next() {

		err := rows.Scan(
			&totalCheckedItems,
			&checkDate,
			&workerCode,
			&productSKU,
			&description,
			&effectiveDate,
			&checkedPrice,
			&pdv,
			&totalOccurrences,
			&productStatus,
			&ticketStatus,
			&ean,
		)
		if err != nil {
			return nil, err
		}

		// Adicionar item à lista
		item := &dto.CheckDetailsItem{
			SKU:           productSKU,
			Description:   description,
			CheckedPrice:  checkedPrice,
			PDV:           pdv,
			EAN:           ean,
			EffectiveDate: effectiveDate,
			ProductStatus: productStatus,
			TicketStatus:  ticketStatus,
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	imageBase64 := base64.StdEncoding.EncodeToString(image)
	result = &dto.CheckDetailsResult{
		TotalCheckedItems: totalCheckedItems,
		TotalOccurrences:  totalOccurrences,
		WorkerID:          fmt.Sprintf("%d", workerCode),
		Date:              checkDate.UTC(),
		Photo:             imageBase64,
		Items:             items,
	}

	if len(items) == 0 {
		return nil, sql.ErrNoRows
	}

	return result, nil
}