package tcg

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.8
The original interface "Screen" can be found in github.com/gdamore/tcell
*/
import (
	"sync/atomic"
	"time"

	tcell "github.com/gdamore/tcell"
	"github.com/gojuno/minimock"

	testify_assert "github.com/stretchr/testify/assert"
)

//ScreenMock implements github.com/gdamore/tcell.Screen
type ScreenMock struct {
	t minimock.Tester

	CanDisplayFunc    func(p rune, p1 bool) (r bool)
	CanDisplayCounter uint64
	CanDisplayMock    mScreenMockCanDisplay

	CharacterSetFunc    func() (r string)
	CharacterSetCounter uint64
	CharacterSetMock    mScreenMockCharacterSet

	ClearFunc    func()
	ClearCounter uint64
	ClearMock    mScreenMockClear

	ColorsFunc    func() (r int)
	ColorsCounter uint64
	ColorsMock    mScreenMockColors

	DisableMouseFunc    func()
	DisableMouseCounter uint64
	DisableMouseMock    mScreenMockDisableMouse

	EnableMouseFunc    func()
	EnableMouseCounter uint64
	EnableMouseMock    mScreenMockEnableMouse

	FillFunc    func(p rune, p1 tcell.Style)
	FillCounter uint64
	FillMock    mScreenMockFill

	FiniFunc    func()
	FiniCounter uint64
	FiniMock    mScreenMockFini

	GetContentFunc    func(p int, p1 int) (r rune, r1 []rune, r2 tcell.Style, r3 int)
	GetContentCounter uint64
	GetContentMock    mScreenMockGetContent

	HasKeyFunc    func(p tcell.Key) (r bool)
	HasKeyCounter uint64
	HasKeyMock    mScreenMockHasKey

	HasMouseFunc    func() (r bool)
	HasMouseCounter uint64
	HasMouseMock    mScreenMockHasMouse

	HideCursorFunc    func()
	HideCursorCounter uint64
	HideCursorMock    mScreenMockHideCursor

	InitFunc    func() (r error)
	InitCounter uint64
	InitMock    mScreenMockInit

	PollEventFunc    func() (r tcell.Event)
	PollEventCounter uint64
	PollEventMock    mScreenMockPollEvent

	PostEventFunc    func(p tcell.Event) (r error)
	PostEventCounter uint64
	PostEventMock    mScreenMockPostEvent

	PostEventWaitFunc    func(p tcell.Event)
	PostEventWaitCounter uint64
	PostEventWaitMock    mScreenMockPostEventWait

	RegisterRuneFallbackFunc    func(p rune, p1 string)
	RegisterRuneFallbackCounter uint64
	RegisterRuneFallbackMock    mScreenMockRegisterRuneFallback

	ResizeFunc    func(p int, p1 int, p2 int, p3 int)
	ResizeCounter uint64
	ResizeMock    mScreenMockResize

	SetCellFunc    func(p int, p1 int, p2 tcell.Style, p3 ...rune)
	SetCellCounter uint64
	SetCellMock    mScreenMockSetCell

	SetContentFunc    func(p int, p1 int, p2 rune, p3 []rune, p4 tcell.Style)
	SetContentCounter uint64
	SetContentMock    mScreenMockSetContent

	SetStyleFunc    func(p tcell.Style)
	SetStyleCounter uint64
	SetStyleMock    mScreenMockSetStyle

	ShowFunc    func()
	ShowCounter uint64
	ShowMock    mScreenMockShow

	ShowCursorFunc    func(p int, p1 int)
	ShowCursorCounter uint64
	ShowCursorMock    mScreenMockShowCursor

	SizeFunc    func() (r int, r1 int)
	SizeCounter uint64
	SizeMock    mScreenMockSize

	SyncFunc    func()
	SyncCounter uint64
	SyncMock    mScreenMockSync

	UnregisterRuneFallbackFunc    func(p rune)
	UnregisterRuneFallbackCounter uint64
	UnregisterRuneFallbackMock    mScreenMockUnregisterRuneFallback
}

//NewScreenMock returns a mock for github.com/gdamore/tcell.Screen
func NewScreenMock(t minimock.Tester) *ScreenMock {
	m := &ScreenMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CanDisplayMock = mScreenMockCanDisplay{mock: m}
	m.CharacterSetMock = mScreenMockCharacterSet{mock: m}
	m.ClearMock = mScreenMockClear{mock: m}
	m.ColorsMock = mScreenMockColors{mock: m}
	m.DisableMouseMock = mScreenMockDisableMouse{mock: m}
	m.EnableMouseMock = mScreenMockEnableMouse{mock: m}
	m.FillMock = mScreenMockFill{mock: m}
	m.FiniMock = mScreenMockFini{mock: m}
	m.GetContentMock = mScreenMockGetContent{mock: m}
	m.HasKeyMock = mScreenMockHasKey{mock: m}
	m.HasMouseMock = mScreenMockHasMouse{mock: m}
	m.HideCursorMock = mScreenMockHideCursor{mock: m}
	m.InitMock = mScreenMockInit{mock: m}
	m.PollEventMock = mScreenMockPollEvent{mock: m}
	m.PostEventMock = mScreenMockPostEvent{mock: m}
	m.PostEventWaitMock = mScreenMockPostEventWait{mock: m}
	m.RegisterRuneFallbackMock = mScreenMockRegisterRuneFallback{mock: m}
	m.ResizeMock = mScreenMockResize{mock: m}
	m.SetCellMock = mScreenMockSetCell{mock: m}
	m.SetContentMock = mScreenMockSetContent{mock: m}
	m.SetStyleMock = mScreenMockSetStyle{mock: m}
	m.ShowMock = mScreenMockShow{mock: m}
	m.ShowCursorMock = mScreenMockShowCursor{mock: m}
	m.SizeMock = mScreenMockSize{mock: m}
	m.SyncMock = mScreenMockSync{mock: m}
	m.UnregisterRuneFallbackMock = mScreenMockUnregisterRuneFallback{mock: m}

	return m
}

type mScreenMockCanDisplay struct {
	mock             *ScreenMock
	mockExpectations *ScreenMockCanDisplayParams
}

//ScreenMockCanDisplayParams represents input parameters of the Screen.CanDisplay
type ScreenMockCanDisplayParams struct {
	p  rune
	p1 bool
}

//Expect sets up expected params for the Screen.CanDisplay
func (m *mScreenMockCanDisplay) Expect(p rune, p1 bool) *mScreenMockCanDisplay {
	m.mockExpectations = &ScreenMockCanDisplayParams{p, p1}
	return m
}

//Return sets up a mock for Screen.CanDisplay to return Return's arguments
func (m *mScreenMockCanDisplay) Return(r bool) *ScreenMock {
	m.mock.CanDisplayFunc = func(p rune, p1 bool) bool {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.CanDisplay method
func (m *mScreenMockCanDisplay) Set(f func(p rune, p1 bool) (r bool)) *ScreenMock {
	m.mock.CanDisplayFunc = f
	return m.mock
}

//CanDisplay implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) CanDisplay(p rune, p1 bool) (r bool) {
	defer atomic.AddUint64(&m.CanDisplayCounter, 1)

	if m.CanDisplayMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.CanDisplayMock.mockExpectations, ScreenMockCanDisplayParams{p, p1},
			"Screen.CanDisplay got unexpected parameters")

		if m.CanDisplayFunc == nil {

			m.t.Fatal("No results are set for the ScreenMock.CanDisplay")

			return
		}
	}

	if m.CanDisplayFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.CanDisplay")
		return
	}

	return m.CanDisplayFunc(p, p1)
}

//CanDisplayMinimockCounter returns a count of Screen.CanDisplay invocations
func (m *ScreenMock) CanDisplayMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.CanDisplayCounter)
}

type mScreenMockCharacterSet struct {
	mock *ScreenMock
}

//Return sets up a mock for Screen.CharacterSet to return Return's arguments
func (m *mScreenMockCharacterSet) Return(r string) *ScreenMock {
	m.mock.CharacterSetFunc = func() string {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.CharacterSet method
func (m *mScreenMockCharacterSet) Set(f func() (r string)) *ScreenMock {
	m.mock.CharacterSetFunc = f
	return m.mock
}

//CharacterSet implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) CharacterSet() (r string) {
	defer atomic.AddUint64(&m.CharacterSetCounter, 1)

	if m.CharacterSetFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.CharacterSet")
		return
	}

	return m.CharacterSetFunc()
}

//CharacterSetMinimockCounter returns a count of Screen.CharacterSet invocations
func (m *ScreenMock) CharacterSetMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.CharacterSetCounter)
}

type mScreenMockClear struct {
	mock *ScreenMock
}

//Return sets up a mock for Screen.Clear to return Return's arguments
func (m *mScreenMockClear) Return() *ScreenMock {
	m.mock.ClearFunc = func() {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.Clear method
func (m *mScreenMockClear) Set(f func()) *ScreenMock {
	m.mock.ClearFunc = f
	return m.mock
}

//Clear implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) Clear() {
	defer atomic.AddUint64(&m.ClearCounter, 1)

	if m.ClearFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.Clear")
		return
	}

	m.ClearFunc()
}

//ClearMinimockCounter returns a count of Screen.Clear invocations
func (m *ScreenMock) ClearMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ClearCounter)
}

type mScreenMockColors struct {
	mock *ScreenMock
}

//Return sets up a mock for Screen.Colors to return Return's arguments
func (m *mScreenMockColors) Return(r int) *ScreenMock {
	m.mock.ColorsFunc = func() int {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.Colors method
func (m *mScreenMockColors) Set(f func() (r int)) *ScreenMock {
	m.mock.ColorsFunc = f
	return m.mock
}

//Colors implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) Colors() (r int) {
	defer atomic.AddUint64(&m.ColorsCounter, 1)

	if m.ColorsFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.Colors")
		return
	}

	return m.ColorsFunc()
}

//ColorsMinimockCounter returns a count of Screen.Colors invocations
func (m *ScreenMock) ColorsMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ColorsCounter)
}

type mScreenMockDisableMouse struct {
	mock *ScreenMock
}

//Return sets up a mock for Screen.DisableMouse to return Return's arguments
func (m *mScreenMockDisableMouse) Return() *ScreenMock {
	m.mock.DisableMouseFunc = func() {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.DisableMouse method
func (m *mScreenMockDisableMouse) Set(f func()) *ScreenMock {
	m.mock.DisableMouseFunc = f
	return m.mock
}

//DisableMouse implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) DisableMouse() {
	defer atomic.AddUint64(&m.DisableMouseCounter, 1)

	if m.DisableMouseFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.DisableMouse")
		return
	}

	m.DisableMouseFunc()
}

//DisableMouseMinimockCounter returns a count of Screen.DisableMouse invocations
func (m *ScreenMock) DisableMouseMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.DisableMouseCounter)
}

type mScreenMockEnableMouse struct {
	mock *ScreenMock
}

//Return sets up a mock for Screen.EnableMouse to return Return's arguments
func (m *mScreenMockEnableMouse) Return() *ScreenMock {
	m.mock.EnableMouseFunc = func() {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.EnableMouse method
func (m *mScreenMockEnableMouse) Set(f func()) *ScreenMock {
	m.mock.EnableMouseFunc = f
	return m.mock
}

//EnableMouse implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) EnableMouse() {
	defer atomic.AddUint64(&m.EnableMouseCounter, 1)

	if m.EnableMouseFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.EnableMouse")
		return
	}

	m.EnableMouseFunc()
}

//EnableMouseMinimockCounter returns a count of Screen.EnableMouse invocations
func (m *ScreenMock) EnableMouseMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.EnableMouseCounter)
}

type mScreenMockFill struct {
	mock             *ScreenMock
	mockExpectations *ScreenMockFillParams
}

//ScreenMockFillParams represents input parameters of the Screen.Fill
type ScreenMockFillParams struct {
	p  rune
	p1 tcell.Style
}

//Expect sets up expected params for the Screen.Fill
func (m *mScreenMockFill) Expect(p rune, p1 tcell.Style) *mScreenMockFill {
	m.mockExpectations = &ScreenMockFillParams{p, p1}
	return m
}

//Return sets up a mock for Screen.Fill to return Return's arguments
func (m *mScreenMockFill) Return() *ScreenMock {
	m.mock.FillFunc = func(p rune, p1 tcell.Style) {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.Fill method
func (m *mScreenMockFill) Set(f func(p rune, p1 tcell.Style)) *ScreenMock {
	m.mock.FillFunc = f
	return m.mock
}

//Fill implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) Fill(p rune, p1 tcell.Style) {
	defer atomic.AddUint64(&m.FillCounter, 1)

	if m.FillMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.FillMock.mockExpectations, ScreenMockFillParams{p, p1},
			"Screen.Fill got unexpected parameters")

		if m.FillFunc == nil {

			m.t.Fatal("No results are set for the ScreenMock.Fill")

			return
		}
	}

	if m.FillFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.Fill")
		return
	}

	m.FillFunc(p, p1)
}

//FillMinimockCounter returns a count of Screen.Fill invocations
func (m *ScreenMock) FillMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.FillCounter)
}

type mScreenMockFini struct {
	mock *ScreenMock
}

//Return sets up a mock for Screen.Fini to return Return's arguments
func (m *mScreenMockFini) Return() *ScreenMock {
	m.mock.FiniFunc = func() {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.Fini method
func (m *mScreenMockFini) Set(f func()) *ScreenMock {
	m.mock.FiniFunc = f
	return m.mock
}

//Fini implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) Fini() {
	defer atomic.AddUint64(&m.FiniCounter, 1)

	if m.FiniFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.Fini")
		return
	}

	m.FiniFunc()
}

//FiniMinimockCounter returns a count of Screen.Fini invocations
func (m *ScreenMock) FiniMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.FiniCounter)
}

type mScreenMockGetContent struct {
	mock             *ScreenMock
	mockExpectations *ScreenMockGetContentParams
}

//ScreenMockGetContentParams represents input parameters of the Screen.GetContent
type ScreenMockGetContentParams struct {
	p  int
	p1 int
}

//Expect sets up expected params for the Screen.GetContent
func (m *mScreenMockGetContent) Expect(p int, p1 int) *mScreenMockGetContent {
	m.mockExpectations = &ScreenMockGetContentParams{p, p1}
	return m
}

//Return sets up a mock for Screen.GetContent to return Return's arguments
func (m *mScreenMockGetContent) Return(r rune, r1 []rune, r2 tcell.Style, r3 int) *ScreenMock {
	m.mock.GetContentFunc = func(p int, p1 int) (rune, []rune, tcell.Style, int) {
		return r, r1, r2, r3
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.GetContent method
func (m *mScreenMockGetContent) Set(f func(p int, p1 int) (r rune, r1 []rune, r2 tcell.Style, r3 int)) *ScreenMock {
	m.mock.GetContentFunc = f
	return m.mock
}

//GetContent implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) GetContent(p int, p1 int) (r rune, r1 []rune, r2 tcell.Style, r3 int) {
	defer atomic.AddUint64(&m.GetContentCounter, 1)

	if m.GetContentMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.GetContentMock.mockExpectations, ScreenMockGetContentParams{p, p1},
			"Screen.GetContent got unexpected parameters")

		if m.GetContentFunc == nil {

			m.t.Fatal("No results are set for the ScreenMock.GetContent")

			return
		}
	}

	if m.GetContentFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.GetContent")
		return
	}

	return m.GetContentFunc(p, p1)
}

//GetContentMinimockCounter returns a count of Screen.GetContent invocations
func (m *ScreenMock) GetContentMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetContentCounter)
}

type mScreenMockHasKey struct {
	mock             *ScreenMock
	mockExpectations *ScreenMockHasKeyParams
}

//ScreenMockHasKeyParams represents input parameters of the Screen.HasKey
type ScreenMockHasKeyParams struct {
	p tcell.Key
}

//Expect sets up expected params for the Screen.HasKey
func (m *mScreenMockHasKey) Expect(p tcell.Key) *mScreenMockHasKey {
	m.mockExpectations = &ScreenMockHasKeyParams{p}
	return m
}

//Return sets up a mock for Screen.HasKey to return Return's arguments
func (m *mScreenMockHasKey) Return(r bool) *ScreenMock {
	m.mock.HasKeyFunc = func(p tcell.Key) bool {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.HasKey method
func (m *mScreenMockHasKey) Set(f func(p tcell.Key) (r bool)) *ScreenMock {
	m.mock.HasKeyFunc = f
	return m.mock
}

//HasKey implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) HasKey(p tcell.Key) (r bool) {
	defer atomic.AddUint64(&m.HasKeyCounter, 1)

	if m.HasKeyMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.HasKeyMock.mockExpectations, ScreenMockHasKeyParams{p},
			"Screen.HasKey got unexpected parameters")

		if m.HasKeyFunc == nil {

			m.t.Fatal("No results are set for the ScreenMock.HasKey")

			return
		}
	}

	if m.HasKeyFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.HasKey")
		return
	}

	return m.HasKeyFunc(p)
}

//HasKeyMinimockCounter returns a count of Screen.HasKey invocations
func (m *ScreenMock) HasKeyMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.HasKeyCounter)
}

type mScreenMockHasMouse struct {
	mock *ScreenMock
}

//Return sets up a mock for Screen.HasMouse to return Return's arguments
func (m *mScreenMockHasMouse) Return(r bool) *ScreenMock {
	m.mock.HasMouseFunc = func() bool {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.HasMouse method
func (m *mScreenMockHasMouse) Set(f func() (r bool)) *ScreenMock {
	m.mock.HasMouseFunc = f
	return m.mock
}

//HasMouse implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) HasMouse() (r bool) {
	defer atomic.AddUint64(&m.HasMouseCounter, 1)

	if m.HasMouseFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.HasMouse")
		return
	}

	return m.HasMouseFunc()
}

//HasMouseMinimockCounter returns a count of Screen.HasMouse invocations
func (m *ScreenMock) HasMouseMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.HasMouseCounter)
}

type mScreenMockHideCursor struct {
	mock *ScreenMock
}

//Return sets up a mock for Screen.HideCursor to return Return's arguments
func (m *mScreenMockHideCursor) Return() *ScreenMock {
	m.mock.HideCursorFunc = func() {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.HideCursor method
func (m *mScreenMockHideCursor) Set(f func()) *ScreenMock {
	m.mock.HideCursorFunc = f
	return m.mock
}

//HideCursor implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) HideCursor() {
	defer atomic.AddUint64(&m.HideCursorCounter, 1)

	if m.HideCursorFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.HideCursor")
		return
	}

	m.HideCursorFunc()
}

//HideCursorMinimockCounter returns a count of Screen.HideCursor invocations
func (m *ScreenMock) HideCursorMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.HideCursorCounter)
}

type mScreenMockInit struct {
	mock *ScreenMock
}

//Return sets up a mock for Screen.Init to return Return's arguments
func (m *mScreenMockInit) Return(r error) *ScreenMock {
	m.mock.InitFunc = func() error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.Init method
func (m *mScreenMockInit) Set(f func() (r error)) *ScreenMock {
	m.mock.InitFunc = f
	return m.mock
}

//Init implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) Init() (r error) {
	defer atomic.AddUint64(&m.InitCounter, 1)

	if m.InitFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.Init")
		return
	}

	return m.InitFunc()
}

//InitMinimockCounter returns a count of Screen.Init invocations
func (m *ScreenMock) InitMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.InitCounter)
}

type mScreenMockPollEvent struct {
	mock *ScreenMock
}

//Return sets up a mock for Screen.PollEvent to return Return's arguments
func (m *mScreenMockPollEvent) Return(r tcell.Event) *ScreenMock {
	m.mock.PollEventFunc = func() tcell.Event {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.PollEvent method
func (m *mScreenMockPollEvent) Set(f func() (r tcell.Event)) *ScreenMock {
	m.mock.PollEventFunc = f
	return m.mock
}

//PollEvent implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) PollEvent() (r tcell.Event) {
	defer atomic.AddUint64(&m.PollEventCounter, 1)

	if m.PollEventFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.PollEvent")
		return
	}

	return m.PollEventFunc()
}

//PollEventMinimockCounter returns a count of Screen.PollEvent invocations
func (m *ScreenMock) PollEventMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.PollEventCounter)
}

type mScreenMockPostEvent struct {
	mock             *ScreenMock
	mockExpectations *ScreenMockPostEventParams
}

//ScreenMockPostEventParams represents input parameters of the Screen.PostEvent
type ScreenMockPostEventParams struct {
	p tcell.Event
}

//Expect sets up expected params for the Screen.PostEvent
func (m *mScreenMockPostEvent) Expect(p tcell.Event) *mScreenMockPostEvent {
	m.mockExpectations = &ScreenMockPostEventParams{p}
	return m
}

//Return sets up a mock for Screen.PostEvent to return Return's arguments
func (m *mScreenMockPostEvent) Return(r error) *ScreenMock {
	m.mock.PostEventFunc = func(p tcell.Event) error {
		return r
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.PostEvent method
func (m *mScreenMockPostEvent) Set(f func(p tcell.Event) (r error)) *ScreenMock {
	m.mock.PostEventFunc = f
	return m.mock
}

//PostEvent implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) PostEvent(p tcell.Event) (r error) {
	defer atomic.AddUint64(&m.PostEventCounter, 1)

	if m.PostEventMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.PostEventMock.mockExpectations, ScreenMockPostEventParams{p},
			"Screen.PostEvent got unexpected parameters")

		if m.PostEventFunc == nil {

			m.t.Fatal("No results are set for the ScreenMock.PostEvent")

			return
		}
	}

	if m.PostEventFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.PostEvent")
		return
	}

	return m.PostEventFunc(p)
}

//PostEventMinimockCounter returns a count of Screen.PostEvent invocations
func (m *ScreenMock) PostEventMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.PostEventCounter)
}

type mScreenMockPostEventWait struct {
	mock             *ScreenMock
	mockExpectations *ScreenMockPostEventWaitParams
}

//ScreenMockPostEventWaitParams represents input parameters of the Screen.PostEventWait
type ScreenMockPostEventWaitParams struct {
	p tcell.Event
}

//Expect sets up expected params for the Screen.PostEventWait
func (m *mScreenMockPostEventWait) Expect(p tcell.Event) *mScreenMockPostEventWait {
	m.mockExpectations = &ScreenMockPostEventWaitParams{p}
	return m
}

//Return sets up a mock for Screen.PostEventWait to return Return's arguments
func (m *mScreenMockPostEventWait) Return() *ScreenMock {
	m.mock.PostEventWaitFunc = func(p tcell.Event) {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.PostEventWait method
func (m *mScreenMockPostEventWait) Set(f func(p tcell.Event)) *ScreenMock {
	m.mock.PostEventWaitFunc = f
	return m.mock
}

//PostEventWait implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) PostEventWait(p tcell.Event) {
	defer atomic.AddUint64(&m.PostEventWaitCounter, 1)

	if m.PostEventWaitMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.PostEventWaitMock.mockExpectations, ScreenMockPostEventWaitParams{p},
			"Screen.PostEventWait got unexpected parameters")

		if m.PostEventWaitFunc == nil {

			m.t.Fatal("No results are set for the ScreenMock.PostEventWait")

			return
		}
	}

	if m.PostEventWaitFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.PostEventWait")
		return
	}

	m.PostEventWaitFunc(p)
}

//PostEventWaitMinimockCounter returns a count of Screen.PostEventWait invocations
func (m *ScreenMock) PostEventWaitMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.PostEventWaitCounter)
}

type mScreenMockRegisterRuneFallback struct {
	mock             *ScreenMock
	mockExpectations *ScreenMockRegisterRuneFallbackParams
}

//ScreenMockRegisterRuneFallbackParams represents input parameters of the Screen.RegisterRuneFallback
type ScreenMockRegisterRuneFallbackParams struct {
	p  rune
	p1 string
}

//Expect sets up expected params for the Screen.RegisterRuneFallback
func (m *mScreenMockRegisterRuneFallback) Expect(p rune, p1 string) *mScreenMockRegisterRuneFallback {
	m.mockExpectations = &ScreenMockRegisterRuneFallbackParams{p, p1}
	return m
}

//Return sets up a mock for Screen.RegisterRuneFallback to return Return's arguments
func (m *mScreenMockRegisterRuneFallback) Return() *ScreenMock {
	m.mock.RegisterRuneFallbackFunc = func(p rune, p1 string) {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.RegisterRuneFallback method
func (m *mScreenMockRegisterRuneFallback) Set(f func(p rune, p1 string)) *ScreenMock {
	m.mock.RegisterRuneFallbackFunc = f
	return m.mock
}

//RegisterRuneFallback implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) RegisterRuneFallback(p rune, p1 string) {
	defer atomic.AddUint64(&m.RegisterRuneFallbackCounter, 1)

	if m.RegisterRuneFallbackMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.RegisterRuneFallbackMock.mockExpectations, ScreenMockRegisterRuneFallbackParams{p, p1},
			"Screen.RegisterRuneFallback got unexpected parameters")

		if m.RegisterRuneFallbackFunc == nil {

			m.t.Fatal("No results are set for the ScreenMock.RegisterRuneFallback")

			return
		}
	}

	if m.RegisterRuneFallbackFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.RegisterRuneFallback")
		return
	}

	m.RegisterRuneFallbackFunc(p, p1)
}

//RegisterRuneFallbackMinimockCounter returns a count of Screen.RegisterRuneFallback invocations
func (m *ScreenMock) RegisterRuneFallbackMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.RegisterRuneFallbackCounter)
}

type mScreenMockResize struct {
	mock             *ScreenMock
	mockExpectations *ScreenMockResizeParams
}

//ScreenMockResizeParams represents input parameters of the Screen.Resize
type ScreenMockResizeParams struct {
	p  int
	p1 int
	p2 int
	p3 int
}

//Expect sets up expected params for the Screen.Resize
func (m *mScreenMockResize) Expect(p int, p1 int, p2 int, p3 int) *mScreenMockResize {
	m.mockExpectations = &ScreenMockResizeParams{p, p1, p2, p3}
	return m
}

//Return sets up a mock for Screen.Resize to return Return's arguments
func (m *mScreenMockResize) Return() *ScreenMock {
	m.mock.ResizeFunc = func(p int, p1 int, p2 int, p3 int) {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.Resize method
func (m *mScreenMockResize) Set(f func(p int, p1 int, p2 int, p3 int)) *ScreenMock {
	m.mock.ResizeFunc = f
	return m.mock
}

//Resize implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) Resize(p int, p1 int, p2 int, p3 int) {
	defer atomic.AddUint64(&m.ResizeCounter, 1)

	if m.ResizeMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.ResizeMock.mockExpectations, ScreenMockResizeParams{p, p1, p2, p3},
			"Screen.Resize got unexpected parameters")

		if m.ResizeFunc == nil {

			m.t.Fatal("No results are set for the ScreenMock.Resize")

			return
		}
	}

	if m.ResizeFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.Resize")
		return
	}

	m.ResizeFunc(p, p1, p2, p3)
}

//ResizeMinimockCounter returns a count of Screen.Resize invocations
func (m *ScreenMock) ResizeMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ResizeCounter)
}

type mScreenMockSetCell struct {
	mock             *ScreenMock
	mockExpectations *ScreenMockSetCellParams
}

//ScreenMockSetCellParams represents input parameters of the Screen.SetCell
type ScreenMockSetCellParams struct {
	p  int
	p1 int
	p2 tcell.Style
	p3 []rune
}

//Expect sets up expected params for the Screen.SetCell
func (m *mScreenMockSetCell) Expect(p int, p1 int, p2 tcell.Style, p3 ...rune) *mScreenMockSetCell {
	m.mockExpectations = &ScreenMockSetCellParams{p, p1, p2, p3}
	return m
}

//Return sets up a mock for Screen.SetCell to return Return's arguments
func (m *mScreenMockSetCell) Return() *ScreenMock {
	m.mock.SetCellFunc = func(p int, p1 int, p2 tcell.Style, p3 ...rune) {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.SetCell method
func (m *mScreenMockSetCell) Set(f func(p int, p1 int, p2 tcell.Style, p3 ...rune)) *ScreenMock {
	m.mock.SetCellFunc = f
	return m.mock
}

//SetCell implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) SetCell(p int, p1 int, p2 tcell.Style, p3 ...rune) {
	defer atomic.AddUint64(&m.SetCellCounter, 1)

	if m.SetCellMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.SetCellMock.mockExpectations, ScreenMockSetCellParams{p, p1, p2, p3},
			"Screen.SetCell got unexpected parameters")

		if m.SetCellFunc == nil {

			m.t.Fatal("No results are set for the ScreenMock.SetCell")

			return
		}
	}

	if m.SetCellFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.SetCell")
		return
	}

	m.SetCellFunc(p, p1, p2, p3...)
}

//SetCellMinimockCounter returns a count of Screen.SetCell invocations
func (m *ScreenMock) SetCellMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SetCellCounter)
}

type mScreenMockSetContent struct {
	mock             *ScreenMock
	mockExpectations *ScreenMockSetContentParams
}

//ScreenMockSetContentParams represents input parameters of the Screen.SetContent
type ScreenMockSetContentParams struct {
	p  int
	p1 int
	p2 rune
	p3 []rune
	p4 tcell.Style
}

//Expect sets up expected params for the Screen.SetContent
func (m *mScreenMockSetContent) Expect(p int, p1 int, p2 rune, p3 []rune, p4 tcell.Style) *mScreenMockSetContent {
	m.mockExpectations = &ScreenMockSetContentParams{p, p1, p2, p3, p4}
	return m
}

//Return sets up a mock for Screen.SetContent to return Return's arguments
func (m *mScreenMockSetContent) Return() *ScreenMock {
	m.mock.SetContentFunc = func(p int, p1 int, p2 rune, p3 []rune, p4 tcell.Style) {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.SetContent method
func (m *mScreenMockSetContent) Set(f func(p int, p1 int, p2 rune, p3 []rune, p4 tcell.Style)) *ScreenMock {
	m.mock.SetContentFunc = f
	return m.mock
}

//SetContent implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) SetContent(p int, p1 int, p2 rune, p3 []rune, p4 tcell.Style) {
	defer atomic.AddUint64(&m.SetContentCounter, 1)

	if m.SetContentMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.SetContentMock.mockExpectations, ScreenMockSetContentParams{p, p1, p2, p3, p4},
			"Screen.SetContent got unexpected parameters")

		if m.SetContentFunc == nil {

			m.t.Fatal("No results are set for the ScreenMock.SetContent")

			return
		}
	}

	if m.SetContentFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.SetContent")
		return
	}

	m.SetContentFunc(p, p1, p2, p3, p4)
}

//SetContentMinimockCounter returns a count of Screen.SetContent invocations
func (m *ScreenMock) SetContentMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SetContentCounter)
}

type mScreenMockSetStyle struct {
	mock             *ScreenMock
	mockExpectations *ScreenMockSetStyleParams
}

//ScreenMockSetStyleParams represents input parameters of the Screen.SetStyle
type ScreenMockSetStyleParams struct {
	p tcell.Style
}

//Expect sets up expected params for the Screen.SetStyle
func (m *mScreenMockSetStyle) Expect(p tcell.Style) *mScreenMockSetStyle {
	m.mockExpectations = &ScreenMockSetStyleParams{p}
	return m
}

//Return sets up a mock for Screen.SetStyle to return Return's arguments
func (m *mScreenMockSetStyle) Return() *ScreenMock {
	m.mock.SetStyleFunc = func(p tcell.Style) {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.SetStyle method
func (m *mScreenMockSetStyle) Set(f func(p tcell.Style)) *ScreenMock {
	m.mock.SetStyleFunc = f
	return m.mock
}

//SetStyle implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) SetStyle(p tcell.Style) {
	defer atomic.AddUint64(&m.SetStyleCounter, 1)

	if m.SetStyleMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.SetStyleMock.mockExpectations, ScreenMockSetStyleParams{p},
			"Screen.SetStyle got unexpected parameters")

		if m.SetStyleFunc == nil {

			m.t.Fatal("No results are set for the ScreenMock.SetStyle")

			return
		}
	}

	if m.SetStyleFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.SetStyle")
		return
	}

	m.SetStyleFunc(p)
}

//SetStyleMinimockCounter returns a count of Screen.SetStyle invocations
func (m *ScreenMock) SetStyleMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SetStyleCounter)
}

type mScreenMockShow struct {
	mock *ScreenMock
}

//Return sets up a mock for Screen.Show to return Return's arguments
func (m *mScreenMockShow) Return() *ScreenMock {
	m.mock.ShowFunc = func() {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.Show method
func (m *mScreenMockShow) Set(f func()) *ScreenMock {
	m.mock.ShowFunc = f
	return m.mock
}

//Show implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) Show() {
	defer atomic.AddUint64(&m.ShowCounter, 1)

	if m.ShowFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.Show")
		return
	}

	m.ShowFunc()
}

//ShowMinimockCounter returns a count of Screen.Show invocations
func (m *ScreenMock) ShowMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ShowCounter)
}

type mScreenMockShowCursor struct {
	mock             *ScreenMock
	mockExpectations *ScreenMockShowCursorParams
}

//ScreenMockShowCursorParams represents input parameters of the Screen.ShowCursor
type ScreenMockShowCursorParams struct {
	p  int
	p1 int
}

//Expect sets up expected params for the Screen.ShowCursor
func (m *mScreenMockShowCursor) Expect(p int, p1 int) *mScreenMockShowCursor {
	m.mockExpectations = &ScreenMockShowCursorParams{p, p1}
	return m
}

//Return sets up a mock for Screen.ShowCursor to return Return's arguments
func (m *mScreenMockShowCursor) Return() *ScreenMock {
	m.mock.ShowCursorFunc = func(p int, p1 int) {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.ShowCursor method
func (m *mScreenMockShowCursor) Set(f func(p int, p1 int)) *ScreenMock {
	m.mock.ShowCursorFunc = f
	return m.mock
}

//ShowCursor implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) ShowCursor(p int, p1 int) {
	defer atomic.AddUint64(&m.ShowCursorCounter, 1)

	if m.ShowCursorMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.ShowCursorMock.mockExpectations, ScreenMockShowCursorParams{p, p1},
			"Screen.ShowCursor got unexpected parameters")

		if m.ShowCursorFunc == nil {

			m.t.Fatal("No results are set for the ScreenMock.ShowCursor")

			return
		}
	}

	if m.ShowCursorFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.ShowCursor")
		return
	}

	m.ShowCursorFunc(p, p1)
}

//ShowCursorMinimockCounter returns a count of Screen.ShowCursor invocations
func (m *ScreenMock) ShowCursorMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.ShowCursorCounter)
}

type mScreenMockSize struct {
	mock *ScreenMock
}

//Return sets up a mock for Screen.Size to return Return's arguments
func (m *mScreenMockSize) Return(r int, r1 int) *ScreenMock {
	m.mock.SizeFunc = func() (int, int) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.Size method
func (m *mScreenMockSize) Set(f func() (r int, r1 int)) *ScreenMock {
	m.mock.SizeFunc = f
	return m.mock
}

//Size implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) Size() (r int, r1 int) {
	defer atomic.AddUint64(&m.SizeCounter, 1)

	if m.SizeFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.Size")
		return
	}

	return m.SizeFunc()
}

//SizeMinimockCounter returns a count of Screen.Size invocations
func (m *ScreenMock) SizeMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SizeCounter)
}

type mScreenMockSync struct {
	mock *ScreenMock
}

//Return sets up a mock for Screen.Sync to return Return's arguments
func (m *mScreenMockSync) Return() *ScreenMock {
	m.mock.SyncFunc = func() {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.Sync method
func (m *mScreenMockSync) Set(f func()) *ScreenMock {
	m.mock.SyncFunc = f
	return m.mock
}

//Sync implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) Sync() {
	defer atomic.AddUint64(&m.SyncCounter, 1)

	if m.SyncFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.Sync")
		return
	}

	m.SyncFunc()
}

//SyncMinimockCounter returns a count of Screen.Sync invocations
func (m *ScreenMock) SyncMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.SyncCounter)
}

type mScreenMockUnregisterRuneFallback struct {
	mock             *ScreenMock
	mockExpectations *ScreenMockUnregisterRuneFallbackParams
}

//ScreenMockUnregisterRuneFallbackParams represents input parameters of the Screen.UnregisterRuneFallback
type ScreenMockUnregisterRuneFallbackParams struct {
	p rune
}

//Expect sets up expected params for the Screen.UnregisterRuneFallback
func (m *mScreenMockUnregisterRuneFallback) Expect(p rune) *mScreenMockUnregisterRuneFallback {
	m.mockExpectations = &ScreenMockUnregisterRuneFallbackParams{p}
	return m
}

//Return sets up a mock for Screen.UnregisterRuneFallback to return Return's arguments
func (m *mScreenMockUnregisterRuneFallback) Return() *ScreenMock {
	m.mock.UnregisterRuneFallbackFunc = func(p rune) {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of Screen.UnregisterRuneFallback method
func (m *mScreenMockUnregisterRuneFallback) Set(f func(p rune)) *ScreenMock {
	m.mock.UnregisterRuneFallbackFunc = f
	return m.mock
}

//UnregisterRuneFallback implements github.com/gdamore/tcell.Screen interface
func (m *ScreenMock) UnregisterRuneFallback(p rune) {
	defer atomic.AddUint64(&m.UnregisterRuneFallbackCounter, 1)

	if m.UnregisterRuneFallbackMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.UnregisterRuneFallbackMock.mockExpectations, ScreenMockUnregisterRuneFallbackParams{p},
			"Screen.UnregisterRuneFallback got unexpected parameters")

		if m.UnregisterRuneFallbackFunc == nil {

			m.t.Fatal("No results are set for the ScreenMock.UnregisterRuneFallback")

			return
		}
	}

	if m.UnregisterRuneFallbackFunc == nil {
		m.t.Fatal("Unexpected call to ScreenMock.UnregisterRuneFallback")
		return
	}

	m.UnregisterRuneFallbackFunc(p)
}

//UnregisterRuneFallbackMinimockCounter returns a count of Screen.UnregisterRuneFallback invocations
func (m *ScreenMock) UnregisterRuneFallbackMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.UnregisterRuneFallbackCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *ScreenMock) ValidateCallCounters() {

	if m.CanDisplayFunc != nil && atomic.LoadUint64(&m.CanDisplayCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.CanDisplay")
	}

	if m.CharacterSetFunc != nil && atomic.LoadUint64(&m.CharacterSetCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.CharacterSet")
	}

	if m.ClearFunc != nil && atomic.LoadUint64(&m.ClearCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.Clear")
	}

	if m.ColorsFunc != nil && atomic.LoadUint64(&m.ColorsCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.Colors")
	}

	if m.DisableMouseFunc != nil && atomic.LoadUint64(&m.DisableMouseCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.DisableMouse")
	}

	if m.EnableMouseFunc != nil && atomic.LoadUint64(&m.EnableMouseCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.EnableMouse")
	}

	if m.FillFunc != nil && atomic.LoadUint64(&m.FillCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.Fill")
	}

	if m.FiniFunc != nil && atomic.LoadUint64(&m.FiniCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.Fini")
	}

	if m.GetContentFunc != nil && atomic.LoadUint64(&m.GetContentCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.GetContent")
	}

	if m.HasKeyFunc != nil && atomic.LoadUint64(&m.HasKeyCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.HasKey")
	}

	if m.HasMouseFunc != nil && atomic.LoadUint64(&m.HasMouseCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.HasMouse")
	}

	if m.HideCursorFunc != nil && atomic.LoadUint64(&m.HideCursorCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.HideCursor")
	}

	if m.InitFunc != nil && atomic.LoadUint64(&m.InitCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.Init")
	}

	if m.PollEventFunc != nil && atomic.LoadUint64(&m.PollEventCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.PollEvent")
	}

	if m.PostEventFunc != nil && atomic.LoadUint64(&m.PostEventCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.PostEvent")
	}

	if m.PostEventWaitFunc != nil && atomic.LoadUint64(&m.PostEventWaitCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.PostEventWait")
	}

	if m.RegisterRuneFallbackFunc != nil && atomic.LoadUint64(&m.RegisterRuneFallbackCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.RegisterRuneFallback")
	}

	if m.ResizeFunc != nil && atomic.LoadUint64(&m.ResizeCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.Resize")
	}

	if m.SetCellFunc != nil && atomic.LoadUint64(&m.SetCellCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.SetCell")
	}

	if m.SetContentFunc != nil && atomic.LoadUint64(&m.SetContentCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.SetContent")
	}

	if m.SetStyleFunc != nil && atomic.LoadUint64(&m.SetStyleCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.SetStyle")
	}

	if m.ShowFunc != nil && atomic.LoadUint64(&m.ShowCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.Show")
	}

	if m.ShowCursorFunc != nil && atomic.LoadUint64(&m.ShowCursorCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.ShowCursor")
	}

	if m.SizeFunc != nil && atomic.LoadUint64(&m.SizeCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.Size")
	}

	if m.SyncFunc != nil && atomic.LoadUint64(&m.SyncCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.Sync")
	}

	if m.UnregisterRuneFallbackFunc != nil && atomic.LoadUint64(&m.UnregisterRuneFallbackCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.UnregisterRuneFallback")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *ScreenMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *ScreenMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *ScreenMock) MinimockFinish() {

	if m.CanDisplayFunc != nil && atomic.LoadUint64(&m.CanDisplayCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.CanDisplay")
	}

	if m.CharacterSetFunc != nil && atomic.LoadUint64(&m.CharacterSetCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.CharacterSet")
	}

	if m.ClearFunc != nil && atomic.LoadUint64(&m.ClearCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.Clear")
	}

	if m.ColorsFunc != nil && atomic.LoadUint64(&m.ColorsCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.Colors")
	}

	if m.DisableMouseFunc != nil && atomic.LoadUint64(&m.DisableMouseCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.DisableMouse")
	}

	if m.EnableMouseFunc != nil && atomic.LoadUint64(&m.EnableMouseCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.EnableMouse")
	}

	if m.FillFunc != nil && atomic.LoadUint64(&m.FillCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.Fill")
	}

	if m.FiniFunc != nil && atomic.LoadUint64(&m.FiniCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.Fini")
	}

	if m.GetContentFunc != nil && atomic.LoadUint64(&m.GetContentCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.GetContent")
	}

	if m.HasKeyFunc != nil && atomic.LoadUint64(&m.HasKeyCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.HasKey")
	}

	if m.HasMouseFunc != nil && atomic.LoadUint64(&m.HasMouseCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.HasMouse")
	}

	if m.HideCursorFunc != nil && atomic.LoadUint64(&m.HideCursorCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.HideCursor")
	}

	if m.InitFunc != nil && atomic.LoadUint64(&m.InitCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.Init")
	}

	if m.PollEventFunc != nil && atomic.LoadUint64(&m.PollEventCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.PollEvent")
	}

	if m.PostEventFunc != nil && atomic.LoadUint64(&m.PostEventCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.PostEvent")
	}

	if m.PostEventWaitFunc != nil && atomic.LoadUint64(&m.PostEventWaitCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.PostEventWait")
	}

	if m.RegisterRuneFallbackFunc != nil && atomic.LoadUint64(&m.RegisterRuneFallbackCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.RegisterRuneFallback")
	}

	if m.ResizeFunc != nil && atomic.LoadUint64(&m.ResizeCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.Resize")
	}

	if m.SetCellFunc != nil && atomic.LoadUint64(&m.SetCellCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.SetCell")
	}

	if m.SetContentFunc != nil && atomic.LoadUint64(&m.SetContentCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.SetContent")
	}

	if m.SetStyleFunc != nil && atomic.LoadUint64(&m.SetStyleCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.SetStyle")
	}

	if m.ShowFunc != nil && atomic.LoadUint64(&m.ShowCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.Show")
	}

	if m.ShowCursorFunc != nil && atomic.LoadUint64(&m.ShowCursorCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.ShowCursor")
	}

	if m.SizeFunc != nil && atomic.LoadUint64(&m.SizeCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.Size")
	}

	if m.SyncFunc != nil && atomic.LoadUint64(&m.SyncCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.Sync")
	}

	if m.UnregisterRuneFallbackFunc != nil && atomic.LoadUint64(&m.UnregisterRuneFallbackCounter) == 0 {
		m.t.Fatal("Expected call to ScreenMock.UnregisterRuneFallback")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *ScreenMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *ScreenMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.CanDisplayFunc == nil || atomic.LoadUint64(&m.CanDisplayCounter) > 0)
		ok = ok && (m.CharacterSetFunc == nil || atomic.LoadUint64(&m.CharacterSetCounter) > 0)
		ok = ok && (m.ClearFunc == nil || atomic.LoadUint64(&m.ClearCounter) > 0)
		ok = ok && (m.ColorsFunc == nil || atomic.LoadUint64(&m.ColorsCounter) > 0)
		ok = ok && (m.DisableMouseFunc == nil || atomic.LoadUint64(&m.DisableMouseCounter) > 0)
		ok = ok && (m.EnableMouseFunc == nil || atomic.LoadUint64(&m.EnableMouseCounter) > 0)
		ok = ok && (m.FillFunc == nil || atomic.LoadUint64(&m.FillCounter) > 0)
		ok = ok && (m.FiniFunc == nil || atomic.LoadUint64(&m.FiniCounter) > 0)
		ok = ok && (m.GetContentFunc == nil || atomic.LoadUint64(&m.GetContentCounter) > 0)
		ok = ok && (m.HasKeyFunc == nil || atomic.LoadUint64(&m.HasKeyCounter) > 0)
		ok = ok && (m.HasMouseFunc == nil || atomic.LoadUint64(&m.HasMouseCounter) > 0)
		ok = ok && (m.HideCursorFunc == nil || atomic.LoadUint64(&m.HideCursorCounter) > 0)
		ok = ok && (m.InitFunc == nil || atomic.LoadUint64(&m.InitCounter) > 0)
		ok = ok && (m.PollEventFunc == nil || atomic.LoadUint64(&m.PollEventCounter) > 0)
		ok = ok && (m.PostEventFunc == nil || atomic.LoadUint64(&m.PostEventCounter) > 0)
		ok = ok && (m.PostEventWaitFunc == nil || atomic.LoadUint64(&m.PostEventWaitCounter) > 0)
		ok = ok && (m.RegisterRuneFallbackFunc == nil || atomic.LoadUint64(&m.RegisterRuneFallbackCounter) > 0)
		ok = ok && (m.ResizeFunc == nil || atomic.LoadUint64(&m.ResizeCounter) > 0)
		ok = ok && (m.SetCellFunc == nil || atomic.LoadUint64(&m.SetCellCounter) > 0)
		ok = ok && (m.SetContentFunc == nil || atomic.LoadUint64(&m.SetContentCounter) > 0)
		ok = ok && (m.SetStyleFunc == nil || atomic.LoadUint64(&m.SetStyleCounter) > 0)
		ok = ok && (m.ShowFunc == nil || atomic.LoadUint64(&m.ShowCounter) > 0)
		ok = ok && (m.ShowCursorFunc == nil || atomic.LoadUint64(&m.ShowCursorCounter) > 0)
		ok = ok && (m.SizeFunc == nil || atomic.LoadUint64(&m.SizeCounter) > 0)
		ok = ok && (m.SyncFunc == nil || atomic.LoadUint64(&m.SyncCounter) > 0)
		ok = ok && (m.UnregisterRuneFallbackFunc == nil || atomic.LoadUint64(&m.UnregisterRuneFallbackCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.CanDisplayFunc != nil && atomic.LoadUint64(&m.CanDisplayCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.CanDisplay")
			}

			if m.CharacterSetFunc != nil && atomic.LoadUint64(&m.CharacterSetCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.CharacterSet")
			}

			if m.ClearFunc != nil && atomic.LoadUint64(&m.ClearCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.Clear")
			}

			if m.ColorsFunc != nil && atomic.LoadUint64(&m.ColorsCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.Colors")
			}

			if m.DisableMouseFunc != nil && atomic.LoadUint64(&m.DisableMouseCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.DisableMouse")
			}

			if m.EnableMouseFunc != nil && atomic.LoadUint64(&m.EnableMouseCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.EnableMouse")
			}

			if m.FillFunc != nil && atomic.LoadUint64(&m.FillCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.Fill")
			}

			if m.FiniFunc != nil && atomic.LoadUint64(&m.FiniCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.Fini")
			}

			if m.GetContentFunc != nil && atomic.LoadUint64(&m.GetContentCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.GetContent")
			}

			if m.HasKeyFunc != nil && atomic.LoadUint64(&m.HasKeyCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.HasKey")
			}

			if m.HasMouseFunc != nil && atomic.LoadUint64(&m.HasMouseCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.HasMouse")
			}

			if m.HideCursorFunc != nil && atomic.LoadUint64(&m.HideCursorCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.HideCursor")
			}

			if m.InitFunc != nil && atomic.LoadUint64(&m.InitCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.Init")
			}

			if m.PollEventFunc != nil && atomic.LoadUint64(&m.PollEventCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.PollEvent")
			}

			if m.PostEventFunc != nil && atomic.LoadUint64(&m.PostEventCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.PostEvent")
			}

			if m.PostEventWaitFunc != nil && atomic.LoadUint64(&m.PostEventWaitCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.PostEventWait")
			}

			if m.RegisterRuneFallbackFunc != nil && atomic.LoadUint64(&m.RegisterRuneFallbackCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.RegisterRuneFallback")
			}

			if m.ResizeFunc != nil && atomic.LoadUint64(&m.ResizeCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.Resize")
			}

			if m.SetCellFunc != nil && atomic.LoadUint64(&m.SetCellCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.SetCell")
			}

			if m.SetContentFunc != nil && atomic.LoadUint64(&m.SetContentCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.SetContent")
			}

			if m.SetStyleFunc != nil && atomic.LoadUint64(&m.SetStyleCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.SetStyle")
			}

			if m.ShowFunc != nil && atomic.LoadUint64(&m.ShowCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.Show")
			}

			if m.ShowCursorFunc != nil && atomic.LoadUint64(&m.ShowCursorCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.ShowCursor")
			}

			if m.SizeFunc != nil && atomic.LoadUint64(&m.SizeCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.Size")
			}

			if m.SyncFunc != nil && atomic.LoadUint64(&m.SyncCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.Sync")
			}

			if m.UnregisterRuneFallbackFunc != nil && atomic.LoadUint64(&m.UnregisterRuneFallbackCounter) == 0 {
				m.t.Error("Expected call to ScreenMock.UnregisterRuneFallback")
			}

			m.t.Fatalf("Some mocks were not called on time: %s", timeout)
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

//AllMocksCalled returns true if all mocked methods were called before the execution of AllMocksCalled,
//it can be used with assert/require, i.e. assert.True(mock.AllMocksCalled())
func (m *ScreenMock) AllMocksCalled() bool {

	if m.CanDisplayFunc != nil && atomic.LoadUint64(&m.CanDisplayCounter) == 0 {
		return false
	}

	if m.CharacterSetFunc != nil && atomic.LoadUint64(&m.CharacterSetCounter) == 0 {
		return false
	}

	if m.ClearFunc != nil && atomic.LoadUint64(&m.ClearCounter) == 0 {
		return false
	}

	if m.ColorsFunc != nil && atomic.LoadUint64(&m.ColorsCounter) == 0 {
		return false
	}

	if m.DisableMouseFunc != nil && atomic.LoadUint64(&m.DisableMouseCounter) == 0 {
		return false
	}

	if m.EnableMouseFunc != nil && atomic.LoadUint64(&m.EnableMouseCounter) == 0 {
		return false
	}

	if m.FillFunc != nil && atomic.LoadUint64(&m.FillCounter) == 0 {
		return false
	}

	if m.FiniFunc != nil && atomic.LoadUint64(&m.FiniCounter) == 0 {
		return false
	}

	if m.GetContentFunc != nil && atomic.LoadUint64(&m.GetContentCounter) == 0 {
		return false
	}

	if m.HasKeyFunc != nil && atomic.LoadUint64(&m.HasKeyCounter) == 0 {
		return false
	}

	if m.HasMouseFunc != nil && atomic.LoadUint64(&m.HasMouseCounter) == 0 {
		return false
	}

	if m.HideCursorFunc != nil && atomic.LoadUint64(&m.HideCursorCounter) == 0 {
		return false
	}

	if m.InitFunc != nil && atomic.LoadUint64(&m.InitCounter) == 0 {
		return false
	}

	if m.PollEventFunc != nil && atomic.LoadUint64(&m.PollEventCounter) == 0 {
		return false
	}

	if m.PostEventFunc != nil && atomic.LoadUint64(&m.PostEventCounter) == 0 {
		return false
	}

	if m.PostEventWaitFunc != nil && atomic.LoadUint64(&m.PostEventWaitCounter) == 0 {
		return false
	}

	if m.RegisterRuneFallbackFunc != nil && atomic.LoadUint64(&m.RegisterRuneFallbackCounter) == 0 {
		return false
	}

	if m.ResizeFunc != nil && atomic.LoadUint64(&m.ResizeCounter) == 0 {
		return false
	}

	if m.SetCellFunc != nil && atomic.LoadUint64(&m.SetCellCounter) == 0 {
		return false
	}

	if m.SetContentFunc != nil && atomic.LoadUint64(&m.SetContentCounter) == 0 {
		return false
	}

	if m.SetStyleFunc != nil && atomic.LoadUint64(&m.SetStyleCounter) == 0 {
		return false
	}

	if m.ShowFunc != nil && atomic.LoadUint64(&m.ShowCounter) == 0 {
		return false
	}

	if m.ShowCursorFunc != nil && atomic.LoadUint64(&m.ShowCursorCounter) == 0 {
		return false
	}

	if m.SizeFunc != nil && atomic.LoadUint64(&m.SizeCounter) == 0 {
		return false
	}

	if m.SyncFunc != nil && atomic.LoadUint64(&m.SyncCounter) == 0 {
		return false
	}

	if m.UnregisterRuneFallbackFunc != nil && atomic.LoadUint64(&m.UnregisterRuneFallbackCounter) == 0 {
		return false
	}

	return true
}
