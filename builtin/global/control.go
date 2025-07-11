package global

const doCode = `
do X Y = do Y
(do X) = X
`

const matchCode = `
(match Case Cases...) X Xs...          = (match' (try Case X Xs...) Cases...) X Xs...
((match) X Xs...)                      = .not-match
((match' (.value Y) Cases...) X Xs...) = Y
(match' .empty Cases...)               = (match Cases...)
`

const retCode = `
(ret X) = X
`

const composeCode = `
((compose) X)       = X
(compose Fs... F) X = (compose Fs...) (F X)
`

const curryCode = `
(curry F S...) X... = F S... X...
`
