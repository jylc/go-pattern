package facade_pattern

import "fmt"

type Amplifier struct {
}

func (a Amplifier) on() {

}
func (a Amplifier) off() {

}
func (a Amplifier) setDvd() {

}

func (a Amplifier) setSurroundSound() {

}

func (a Amplifier) setVolume(num int) {

}

type Tuner struct {
}

type DvdPlayer struct {
}

func (dvd DvdPlayer) on() {

}

func (dvd DvdPlayer) off() {

}
func (dvd DvdPlayer) play(movie string) {

}
func (dvd DvdPlayer) stop() {

}

func (dvd DvdPlayer) eject() {

}

type CdPlayer struct {
}

type Projector struct {
}

func (p Projector) on() {

}

func (p Projector) off() {

}

func (p Projector) wideScreenMode() {

}

type TheaterLights struct {
}

func (light TheaterLights) dim(num int) {

}
func (light TheaterLights) on() {

}

type Screen struct {
}

func (s Screen) down() {

}

func (s Screen) up() {

}

type PopcornPopper struct {
}

func (p PopcornPopper) on() {

}
func (p PopcornPopper) off() {

}

func (p PopcornPopper) pop() {

}

type HomeTheaterFacade struct {
	amp       Amplifier
	tuner     Tuner
	dvd       DvdPlayer
	cd        CdPlayer
	projector Projector
	lights    TheaterLights
	screen    Screen
	popper    PopcornPopper
}

func NewHomeTheaterFacade(amp Amplifier,
	tuner Tuner,
	dvd DvdPlayer,
	cd CdPlayer,
	projector Projector,
	lights TheaterLights,
	screen Screen,
	popper PopcornPopper) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		amp:       amp,
		tuner:     tuner,
		dvd:       dvd,
		cd:        cd,
		projector: projector,
		lights:    lights,
		screen:    screen,
		popper:    popper,
	}
}

func (h HomeTheaterFacade) watchMovie(movie string) {
	fmt.Println("Get ready to watch a movie...")
	h.popper.on()
	h.popper.pop()
	h.lights.dim(10)
	h.screen.down()
	h.projector.on()
	h.projector.wideScreenMode()
	h.amp.on()
	h.amp.setDvd()
	h.amp.setSurroundSound()
	h.amp.setVolume(5)
	h.dvd.on()
	h.dvd.play(movie)
}
func (h HomeTheaterFacade) endMovie() {
	fmt.Println("Shutting movie theater down...")
	h.popper.off()
	h.lights.on()
	h.screen.up()
	h.projector.off()
	h.amp.off()
	h.dvd.stop()
	h.dvd.eject()
	h.dvd.off()
}
