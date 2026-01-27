module Drawing exposing (display)

import Svg exposing (Svg, svg, line)
import Svg.Attributes exposing (..)
import Basics exposing (pi, cos, sin)
import ParserTcTurtle exposing (TurtleProgram, Instr(..))


type alias State =
    { x : Float
    , y : Float
    , angle : Float
    , lines : List Line
    }


type alias Line =
    { x1 : Float
    , y1 : Float
    , x2 : Float
    , y2 : Float
    }


step : Instr -> State -> State
step instr state =
    case instr of
        Forward d ->
            let
                rad = state.angle * (pi / 180)
                newX = state.x + toFloat d * cos rad
                newY = state.y + toFloat d * sin rad
                newLine = { x1 = state.x, y1 = state.y, x2 = newX, y2 = newY }
            in
            { state | x = newX, y = newY, lines = newLine :: state.lines }

        Left a ->
            { state | angle = state.angle + toFloat a }

        Right a ->
            { state | angle = state.angle - toFloat a }

        Repeat n instrs ->
            List.foldl (\_ acc -> execute instrs acc) state (List.repeat n ())


execute : TurtleProgram -> State -> State
execute instrs state =
    List.foldl step state instrs


display : TurtleProgram -> Svg msg
display instrs =
    let
        initialState =
            { x = 250
            , y = 250
            , angle = 0
            , lines = []
            }

        finalState =
            execute instrs initialState
    in
    svg
        [ width "500", height "500", viewBox "0 0 500 500" ]
        (List.map lineToSvg (List.reverse finalState.lines))


lineToSvg : Line -> Svg msg
lineToSvg l =
    line
        [ x1 (String.fromFloat l.x1)
        , y1 (String.fromFloat l.y1)
        , x2 (String.fromFloat l.x2)
        , y2 (String.fromFloat l.y2)
        , stroke "black"
        , strokeWidth "1"
        ]
        []
