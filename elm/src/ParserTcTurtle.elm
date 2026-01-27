module ParserTcTurtle exposing (Instr(..), TurtleProgram, read)

import Parser exposing (Parser, DeadEnd, (|.), (|=))
import Parser as P


-- TYPES

type Instr
    = Forward Int
    | Left Int
    | Right Int
    | Repeat Int (List Instr)


type alias TurtleProgram =
    List Instr


-- TOP-LEVEL PARSER

read : String -> Result (List DeadEnd) TurtleProgram
read str =
    P.run programParser str


-- PARSERS

programParser : Parser TurtleProgram
programParser =
    P.succeed identity
        |. P.spaces
        |. P.symbol "["
        |= instructionList
        |. P.symbol "]"


instructionList : Parser (List Instr)
instructionList =
    P.sequence
        { start = ""
        , separator = ","
        , end = ""
        , spaces = P.spaces
        , item = instruction
        , trailing = P.Forbidden
        }


instruction : Parser Instr
instruction =
    P.lazy <| \_ ->
        P.oneOf
            [ forwardParser
            , leftParser
            , rightParser
            , repeatParser
            ]


forwardParser : Parser Instr
forwardParser =
    P.succeed Forward
        |. P.keyword "Forward"
        |. P.spaces
        |= P.int


leftParser : Parser Instr
leftParser =
    P.succeed Left
        |. P.keyword "Left"
        |. P.spaces
        |= P.int


rightParser : Parser Instr
rightParser =
    P.succeed Right
        |. P.keyword "Right"
        |. P.spaces
        |= P.int


repeatParser : Parser Instr
repeatParser =
    P.succeed Repeat
        |. P.keyword "Repeat"
        |. P.spaces
        |= P.int
        |. P.spaces
        |. P.symbol "["
        |= instructionList
        |. P.symbol "]"
