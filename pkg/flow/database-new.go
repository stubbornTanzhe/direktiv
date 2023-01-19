package flow

import (
	"context"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

type CacheData struct {
	Namespace *Namespace
	Inodes    []*Inode
	Workflow  *Workflow
	Ref       *Ref
	Revision  *Revision
	Instance  *Instance
}

func (cached *CacheData) Path() string {
	var elems []string
	for _, ino := range cached.Inodes {
		elems = append(elems, ino.Name)
	}
	return strings.Join(elems, "/")
}

func (cached *CacheData) Dir() string {
	return filepath.Dir(cached.Path())
}

func (cached *CacheData) Reset() {
	cached.Inodes = make([]*Inode, 0)
}

func (cached *CacheData) Inode() *Inode {
	return cached.Inodes[len(cached.Inodes)-1]
}

func (cached *CacheData) ParentInode() *Inode {
	return cached.Inodes[len(cached.Inodes)-2]
}

type Namespace struct {
	ID        uuid.UUID `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Config    string    `json:"config,omitempty"`
	Name      string    `json:"name,omitempty"`
	Root      uuid.UUID `json:"root,omitempty"`
}

type Inode struct {
	ID           uuid.UUID `json:"id,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	Name         string    `json:"name,omitempty"`
	Type         string    `json:"type,omitempty"`
	Attributes   []string  `json:"attributes,omitempty"`
	ExtendedType string    `json:"expandedType,omitempty"`
	ReadOnly     bool      `json:"readOnly,omitempty"`
	Namespace    uuid.UUID `json:"namespace,omitempty"`
	Children     []*Inode  `json:"children,omitempty"`
	Parent       uuid.UUID `json:"parent,omitempty"`
	Workflow     uuid.UUID `json:"workflow,omitempty"`
	Mirror       uuid.UUID `json:"mirror,omitempty"`
}

type Workflow struct {
	ID          uuid.UUID   `json:"id,omitempty"`
	Live        bool        `json:"live,omitempty"`
	LogToEvents string      `json:"logToEvents,omitempty"`
	ReadOnly    bool        `json:"readOnly,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty"`
	Namespace   uuid.UUID   `json:"namespace,omitempty"`
	Inode       uuid.UUID   `json:"inode,omitempty"`
	Refs        []*Ref      `json:"refs,omitempty"`
	Revisions   []*Revision `json:"revision,omitempty"`
	Routes      []*Route    `json:"route,omitempty"`
}

type Ref struct {
	ID        uuid.UUID `json:"id"`
	Immutable bool      `json:"immutable,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Revision  uuid.UUID `json:"revision,omitempty"`
}

type Revision struct {
	ID        uuid.UUID              `json:"id"`
	CreatedAt time.Time              `json:"created_at,omitempty"`
	Hash      string                 `json:"hash,omitempty"`
	Source    []byte                 `json:"source,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Workflow  uuid.UUID              `json:"workflow,omitempty"`
}

type Route struct {
	ID     uuid.UUID `json:"id"`
	Weight int       `json:"weight,omitempty"`
	Ref    *Ref      `json:"ref,omitempty"`
}

type Instance struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	EndAt        time.Time `json:"end_at,omitempty"`
	Status       string    `json:"status,omitempty"`
	As           string    `json:"as,omitempty"`
	ErrorCode    string    `json:"errorCode,omitempty"`
	ErrorMessage string    `json:"errorMessage,omitempty"`
	Invoker      string    `json:"invoker,omitempty"`
	Namespace    uuid.UUID `json:"namespace,omitempty"`
	Workflow     uuid.UUID `json:"workflow,omitempty"`
	Revision     uuid.UUID `json:"revision,omitempty"`
	Runtime      uuid.UUID `json:"runtime,omitempty"`
}

type InstanceRuntime struct {
	ID              uuid.UUID `json:"id"`
	Input           []byte    `json:"input,omitempty"`
	Data            string    `json:"data,omitempty"`
	Controller      string    `json:"controller,omitempty"`
	Memory          string    `json:"memory,omitempty"`
	Flow            []string  `json:"flow,omitempty"`
	Output          string    `json:"output,omitempty"`
	StateBeginTime  time.Time `json:"stateBeginTime,omitempty"`
	Deadline        time.Time `json:"deadline,omitempty"`
	Attempts        int       `json:"attempts,omitempty"`
	CallerData      string    `json:"caller_data,omitempty"`
	InstanceContext string    `json:"instanceContext,omitempty"`
	StateContext    string    `json:"stateContext,omitempty"`
	Metadata        string    `json:"metadata,omitempty"`
	Caller          uuid.UUID `json:"caller,omitempty"`
}

type Annotation struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Size      int       `json:"size,omitempty"`
	Hash      string    `json:"checksum"`
	Data      []byte    `json:"data,omitempty"`
	MimeType  string    `json:"mime_type,omitempty"`
}

type VarRef struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name,omitempty"`
	Behaviour string    `json:"behaviour,omitempty"`
	VarData   uuid.UUID `json:"vardata,omitempty"`
}

type VarData struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Size      int       `json:"size,omitempty"`
	Hash      string    `json:"hash,omitempty"`
	Data      []byte    `json:"data,omitempty"`
	MimeType  string    `json:"mime_type,omitempty"`
	RefCount  int       `json:"ref_count,omitempty"`
}

type Mirror struct {
	ID         uuid.UUID  `json:"id,omitempty"`
	URL        string     `json:"url,omitempty"`
	Ref        string     `json:"ref,omitempty"`
	Cron       string     `json:"cron,omitempty"`
	PublicKey  string     `json:"publicKey,omitempty"`
	PrivateKey string     `json:"private_key,omitempty"`
	Passphrase string     `json:"passphrase,omitempty"`
	Commit     string     `json:"commit,omitempty"`
	LastSync   *time.Time `json:"last_sync,omitempty"`
	UpdatedAt  time.Time  `json:"updated_at,omitempty"`
}

type MirrorActivity struct {
	ID         uuid.UUID `json:"id,omitempty"`
	Type       string    `json:"type,omitempty"`
	Status     string    `json:"status,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	EndAt      time.Time `json:"end_at,omitempty"`
	Controller string    `json:"controller,omitempty"`
	Deadline   time.Time `json:"deadline,omitempty"`
}

type Transaction interface{}

type Database interface {
	Namespace(ctx context.Context, tx Transaction, id uuid.UUID) (*Namespace, error)
	NamespaceByName(ctx context.Context, tx Transaction, namespace string) (*Namespace, error)
	Inode(ctx context.Context, tx Transaction, id uuid.UUID) (*Inode, error)
	Workflow(ctx context.Context, tx Transaction, id uuid.UUID) (*Workflow, error)
	Revision(ctx context.Context, tx Transaction, id uuid.UUID) (*Revision, error)
	Instance(ctx context.Context, tx Transaction, id uuid.UUID) (*Instance, error)
	InstanceRuntime(ctx context.Context, tx Transaction, id uuid.UUID) (*InstanceRuntime, error)
	NamespaceAnnotation(ctx context.Context, tx Transaction, nsID uuid.UUID, key string) (*Annotation, error)
	InodeAnnotation(ctx context.Context, tx Transaction, inodeID uuid.UUID, key string) (*Annotation, error)
	WorkflowAnnotation(ctx context.Context, tx Transaction, wfID uuid.UUID, key string) (*Annotation, error)
	InstanceAnnotation(ctx context.Context, tx Transaction, instID uuid.UUID, key string) (*Annotation, error)
	ThreadVariables(ctx context.Context, tx Transaction, instID uuid.UUID) ([]*VarRef, error)
	NamespaceVariableRef(ctx context.Context, tx Transaction, nsID uuid.UUID, key string) (*VarRef, error)
	WorkflowVariableRef(ctx context.Context, tx Transaction, wfID uuid.UUID, key string) (*VarRef, error)
	InstanceVariableRef(ctx context.Context, tx Transaction, instID uuid.UUID, key string) (*VarRef, error)
	ThreadVariableRef(ctx context.Context, tx Transaction, instID uuid.UUID, key string) (*VarRef, error)
	VariableData(ctx context.Context, tx Transaction, id uuid.UUID, load bool) (*VarData, error)
	Mirror(ctx context.Context, tx Transaction, id uuid.UUID) (*Mirror, error)
}
