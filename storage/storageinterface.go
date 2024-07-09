package storage
// CRUDRepository defines the interface for CRUD operations
type CRUDRepository interface {
    // Create creates a new record
    Create(interface{}) error
    // Read retrieves a record by ID
    Read(uint) (interface{}, error)
    // Update updates an existing record
    Update(interface{}) error
    // Delete deletes a record by ID
    Delete(uint) error
	// ReadAll retrieves all records
    ReadAll() ([]interface{}, error)
}
