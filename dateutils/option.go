package dateutils

// Option is conf option.
type Option func(o *options)

type WeekMode int

const ()

type options struct {
	weekMode WeekMode
}

// WithWeekMode is conf WeekMode
func WithWeekMode(weekMode WeekMode) Option {
	return Option(func(o *options) {
		o.weekMode = weekMode
	})
}
