package global

const mapCode = `
map F A                        = (map' F ($)) A
((map' _ B) (_))               = B
(map' F ($ Ys...)) (T X Xs...) = (map' F ($ Ys... (F X))) (T Xs...)
`

const filterCode = `
filter P A                        = (filter' P ($)) A
((filter' _ B) (_))               = B
(filter' P ($ Ys...)) (T X Xs...) = (match
  \true  [(filter' P ($ Ys... X)) (T Xs...)]
  \false [(filter' P ($ Ys...)  ) (T Xs...)]
) (P X)
`
