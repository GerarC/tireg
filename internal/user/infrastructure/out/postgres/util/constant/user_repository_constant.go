package constant

const (
	UsernameUniqueConstraint = "users_username_key"
	EmailUniqueConstraint    = "users_email_key"

	InsertUserQuery = `INSERT INTO users (first_name, last_name, username, email, password_hash, created_by, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, first_name, last_name, username, email, password_hash, created_at, created_by, updated_at, updated_by`

	ExistsByUsernameQuery = `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`
	ExistsByEmailQuery    = `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	SelectByIdentifierQuery = `SELECT id, first_name, last_name, username, email, password_hash, created_at, created_by, updated_at, updated_by FROM users WHERE username = $1 OR email = $1`
)
