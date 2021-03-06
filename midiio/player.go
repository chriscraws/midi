package midiio

/*
import (
	"bytes"
	"fmt"
	"github.com/gomidi/midi"
	"github.com/gomidi/midi/live/midiwriter"
	"github.com/gomidi/midi/messages/channel"
	"github.com/gomidi/midi/messages/meta"
	"github.com/gomidi/midi/smf"
	"io"
	"time"
)

// NewPlayer plays from a single track to reader
// src is expected to be of format SMF0. Use smftrack to convert from SMF1 to SMF0.
func NewPlayer(src smf.Reader) (io.Reader, error) {
	err := src.ReadHeader()

	if err != nil {
		return nil, err
	}

	if src.Header().Format != smf.SMF0 {
		return nil, fmt.Errorf("only SMF0 files supported, sorry, please convert your file first")
	}

	ti, isMetric := src.Header().TimeFormat.(smf.MetricTicks)
	if !isMetric {
		return nil, fmt.Errorf("only metric timeformat supported, sorry")
	}

	p := &smfPlayer{}
	p.rd = src
	p.metricTicks = ti
	p.tempo = 120 // default
	p.to = midiwriter.New(&p.bf)
	return p, nil
}

// it is suggested to make the data buffer at least 3 bytes long
// sysex are skipped
func (p *smfPlayer) Read(data []byte) (n int, err error) {
	if len(data) < 3 {
		return 0, fmt.Errorf("data buffer too small must be 3 bytes or more")
	}

	msg, err := p.rd.Read()

	if err != nil {
		return
	}

	delta := p.rd.Delta()

	time.Sleep(p.metricTicks.TempoDuration(p.tempo, delta))

	// track ended
	if msg == meta.EndOfTrack {
		return 0, io.EOF

	}

	switch v := msg.(type) {

	case meta.Tempo:
		p.tempo = v.BPM()

	default:
		if cm, ok := msg.(channel.Message); ok {
			_, err = p.to.Write(cm)

			if err != nil {
				return
			}
		}

		// skip sysex messages, since they are too long
		// skip meta messages, since they can't be used live
	}

	return p.bf.Read(data)
}
*/

/*
// roundFloat rounds the given float by the given decimals after the dot
func roundFloat(x float64, decimals int) float64 {
	// return roundFloat(x, numDig(x)+decimals)
	frep := strconv.FormatFloat(x, 'f', decimals, 64)
	f, _ := strconv.ParseFloat(frep, 64)
	return f
}
*/

/*
 */
/*
type smfPlayer struct {
	metricTicks smf.MetricTicks
	tempo       uint32
	bf          bytes.Buffer
	to          midi.Writer
	rd          smf.Reader
}
*/
