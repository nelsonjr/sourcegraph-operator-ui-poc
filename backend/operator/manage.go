package operator

type K8sManager interface {
	Status() *status
	Install(version string) error
}

func New() K8sManager {
	return &manager{}
}

type status struct {
	Stage      Stage   `json:"stage"`
	Version    *string `json:"version"` // current version, nil if not installed
	Deployment struct {
		Version string   `json:"version"` // version being installed/upgraded
		Errors  []string `json:"errors"`
	} `json:"deployment"`
	Tasks []Task `json:"tasks"`
}

type Stage string

const (
	StageUnknown         Stage = "unknown"
	StageIdle            Stage = "idle"
	StageInstalling      Stage = "installing"
	StageUpgrading       Stage = "upgrading"
	StageWaitingForAdmin Stage = "waitingForAdmin"
	StageRefresh         Stage = "refresh"
)

type manager struct{}

// Asks the Operator to kick off a new installation of the specified version.
//
// Returns an error if the installation was not successful,
// if the version is not supported, or a version is already installed.
//
// Once the request is accepted, the status can be tracked via the Status() method.
func (*manager) Install(version string) error {
	panic("unimplemented")
}

// Returns the current status of the Operator.
func (*manager) Status() *status {
	panic("unimplemented")
}
