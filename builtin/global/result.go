package global

const okCode = `
((ok X) .value)  = X
((ok X) .error?) = false
((ok X) .ok?)    = true
`

const errorCode = `
((error Msg) .error)  = Msg
((error Msg) .error?) = true
((error Msg) .ok?)    = false
`
