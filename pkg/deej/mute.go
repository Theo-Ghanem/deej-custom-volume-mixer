package deej

// MuteSession is a compatibility wrapper that mutes mapped sessions for the given slider.
// It delegates to sessionMap.MuteSessionsForSlider which implements the actual logic.
func (d *Deej) MuteSession(sliderID int) {
	if d.sessions == nil {
		return
	}

	d.sessions.MuteSessionsForSlider(sliderID)
}