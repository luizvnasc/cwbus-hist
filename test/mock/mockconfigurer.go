package mock

import "github.com/luizvnasc/cwbus-hist/config"

// MockConfigurer é um confgiurer para ser utilizado em testes.
type MockConfigurer struct {
	serviceURL string
	urbsCode   string
	dbName     string
	dbStrConn  string
	wakeUpURL  string
}

// SetServiceURL é o setter do serviceURL
func (mc *MockConfigurer) SetServiceURL(url string) {
	mc.serviceURL = url
}

// ServiceURL getter
func (mc MockConfigurer) ServiceURL() string {
	return mc.serviceURL
}

// SetUrbsCode é o setter do urbsCode
func (mc *MockConfigurer) SetUrbsCode(code string) {
	mc.urbsCode = code
}

// UrbsCode getter
func (mc MockConfigurer) UrbsCode() string {
	return mc.urbsCode
}

// SetDBName é o setter do dbName
func (mc *MockConfigurer) SetDBName(name string) {
	mc.dbName = name
}

// DBName getter
func (mc MockConfigurer) DBName() string {
	return mc.dbName
}

// SetDBStrConn é o setter do dbStrConn
func (mc *MockConfigurer) SetDBStrConn(str string) {
	mc.dbStrConn = str
}

// DBStrConn getter
func (mc MockConfigurer) DBStrConn() string {
	return mc.dbStrConn
}

// SetWakeUpURL é o setter do wakeupURL
func (mc *MockConfigurer) SetWakeUpURL(url string) {
	mc.wakeUpURL = url
}

// WakeUpURL getter
func (mc MockConfigurer) WakeUpURL() string {
	return mc.wakeUpURL
}

// ConfigToMock cria um MockConfigurer a partir de um Configurer
func ConfigToMock(c config.Configurer) MockConfigurer {
	return MockConfigurer{
		serviceURL: c.ServiceURL(),
		urbsCode:   c.UrbsCode(),
		dbName:     c.DBName(),
		dbStrConn:  c.DBStrConn(),
		wakeUpURL:  c.WakeUpURL(),
	}
}
