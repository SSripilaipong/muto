main = game (init-state)

game S = (match
    \(.end R)    [display-result (S .board) R]
    \(.continue) [game (next-round S)]
) (evaluate-result S)

next-round S = do
    (print-board! (S .board))
    (print! (string-concat "player" (string (S .current-player)) "input:"))
    ((match
        \(.valid ($ R C)) [(S (.play R C)) .swap-players]
        \.invalid         [do (print! "invalid input") S]
    ) (validate-input (S .board) (input! $)))

evaluate-result S = (match
    \(.value X) [ret X]
    \.empty     [(match
        \true     [.end .draw]
        \false    [.continue]
    ) ((S .board) .is-full?)]
) (winner (S .board))

winner Board = first-match
    \(.value P) [.end (.win P)]
    (map same-value (win-cases Board))
win-cases Board = append-all
    (Board .rows) (Board .columns) (Board .diags)

same-value = (match
    \(_ V)         [.value V]
    \(T V V Vs...) [same-value (T V Vs...)]
    \(T _...)      [ret .empty]
)


init-state = state (players "o" "x") (empty-board)
(state P B) M = (match
    \.players        [ret P]
    \.current-player [P .current]
    \.swap-players   [state (P .swap) B]
    \.board          [ret B]
    \(.play R C)     [state P (B (.set R C (P .current)))]
) M
(players X Y) M = (match
    \.current [ret X]
    \.swap    [players Y X]
) M


validate-input Board Input = (match
    \.unknown [ret .invalid]
    \X        [Board (.validate-availability (num-to-pos X))]
) (string-to-num Input)
num-to-pos X = (with (- X 1)) \Y [
    $ (+ 1 (div Y 3)) (+ 1 (mod Y 3))
]


empty-board = board (arr (arr 1 2 3) (arr 4 5 6) (arr 7 8 9))
(board Arr) M = (with (board Arr)) \Self [
    (match
        \(.row N)   [(use row) (Arr (.get N))]
        \(.col N)   [(use col) (map \X[X (.get N)] Arr)]
        \.rows      [map (use row) Arr]
        \.columns   [(use arr) (map \N[Self (.col N)] (count (Arr .length)))]
        \(.get R C) [(Self (.row R)) (.get C)]

        \(.set-row N Rx) [board (Arr (.set N Rx))]
        \(.set R C X)    [
            (with ((Self (.row R)) (.set C X))) \NewR [
                Self (.set-row R NewR)
            ]
        ]

        \.diags [(with (Arr .length)) \Length [
            arr ((use arr) (map \N[Self (.get N N)] (count Length)))
                ((use arr) (map \N[Self (.get N (+ 1 (- Length N)))] (count Length)))
        ]]
        \(.validate-availability ($ R C)) [(match
            \"o" [ret .invalid]
            \"x" [ret .invalid]
            \_   [.valid ($ R C)]
        ) (Self (.get R C))]
        \.is-full? [== 0
            ((use sum) (map (match \"o"[ret 0] \"x"[ret 0] \_[ret 1]) ((use append-all) Arr)))
        ]
    ) M
]

(row S...) X Xs... = (arr S...) X Xs...
(col S...) X Xs... = (arr S...) X Xs...

display-result B M = do
    (print-board! B)
    (print! ((match
        \(.win X) [string-concat "winner is" (string X)]
        \.draw    [ret "draw"]
    ) M))

print-board! B = map print-row! (B .rows)
print-row! R = print! ((use string-concat) (map string R))

string-to-num X = (match
    \"1" [ret 1]
    \"2" [ret 2]
    \"3" [ret 3]
    \"4" [ret 4]
    \"5" [ret 5]
    \"6" [ret 6]
    \"7" [ret 7]
    \"8" [ret 8]
    \"9" [ret 9]
    \_   [ret .unknown]
) X










string-concat X Y = string-concat (++ (++ X " ") Y)
string-concat X = ret X

first-match _ ($) = ret .empty
first-match P ($ X Xs...) = (match
    \(.value R) [.value R]
    \.empty   [first-match P ($ Xs...)]
) (try P X)

(arr S...) M = (match
    \($ X Xs...) (.get N) [(match
        \1 [ret X]
        \_ [(arr Xs...) (.get (- N 1))]
    ) N]

    \($ X Xs...) (.set N Y) [(match
        \1 [arr Y Xs...]
        \_ [(arr X) (.append ((arr Xs...) (.set (- N 1) Y)))]
    ) N]

    \($ Xs...) (.append (arr Ys...)) [
        arr Xs... Ys...
    ]

    \($)         .length [ret 0]
    \($ X Xs...) .length [+ 1 ((arr Xs...) .length)]
) ($ S...) M

count 1 = $ 1
count N = \($ Rest...)[$ Rest... N] (count (- N 1))

append-all (T Xs...) (_ Ys...) = append-all (T Xs... Ys...)
append-all R = ret R

sum X Y = sum (+ X Y)
sum X = ret X
