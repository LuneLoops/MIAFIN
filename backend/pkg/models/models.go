package models

import (
	"time"

	"gorm.io/gorm"
)

// Usuario representa un usuario del sistema
type Usuario struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	Nombre       string         `gorm:"not null" json:"nombre"`
	Email        string         `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string         `gorm:"not null" json:"-"`
	Rol          string         `gorm:"not null;default:'asesor'" json:"rol"` // admin, asesor
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// Cliente representa un cliente del sistema
type Cliente struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Nombre    string         `gorm:"not null" json:"nombre"`
	CI        string         `gorm:"uniqueIndex;not null" json:"ci"`
	Telefono  string         `json:"telefono"`
	GrupoID   *uint          `json:"grupo_id"`
	Grupo     *Grupo         `gorm:"foreignKey:GrupoID" json:"grupo,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Grupo representa un grupo comunal
type Grupo struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Nombre      string         `gorm:"not null" json:"nombre"`
	Descripcion string         `json:"descripcion"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// Prestamo representa un préstamo
type Prestamo struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Tipo        string         `gorm:"not null" json:"tipo"` // comunal, individual
	ClienteID   *uint          `json:"cliente_id"`
	Cliente     *Cliente       `gorm:"foreignKey:ClienteID" json:"cliente,omitempty"`
	GrupoID     *uint          `json:"grupo_id"`
	Grupo       *Grupo         `gorm:"foreignKey:GrupoID" json:"grupo,omitempty"`
	Monto       float64        `gorm:"not null" json:"monto"`
	TasaInteres float64        `gorm:"not null" json:"tasa_interes"`
	PlazoMeses  int            `gorm:"not null" json:"plazo_meses"`
	SaldoActual float64        `gorm:"not null" json:"saldo_actual"`
	Estado      string         `gorm:"not null;default:'aprobado'" json:"estado"` // aprobado, desembolsado, pagado
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// Pago representa un pago de préstamo
type Pago struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	PrestamoID uint           `gorm:"not null" json:"prestamo_id"`
	Prestamo   *Prestamo      `gorm:"foreignKey:PrestamoID" json:"prestamo,omitempty"`
	Monto      float64        `gorm:"not null" json:"monto"`
	FechaPago  time.Time      `gorm:"not null" json:"fecha_pago"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

// LogTransaccion representa el registro de auditoría
type LogTransaccion struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	UsuarioID   uint      `gorm:"not null" json:"usuario_id"`
	Usuario     *Usuario  `gorm:"foreignKey:UsuarioID" json:"usuario,omitempty"`
	Accion      string    `gorm:"not null" json:"accion"`
	Entidad     string    `gorm:"not null" json:"entidad"`
	EntidadID   uint      `json:"entidad_id"`
	Descripcion string    `json:"descripcion"`
	CreatedAt   time.Time `json:"created_at"`
}
