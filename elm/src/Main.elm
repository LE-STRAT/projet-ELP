module Main exposing (main)

import Browser

import Html exposing (Html, div, text, textarea, button, h2)
import Html.Attributes exposing (value, style, rows, cols)
import Html.Events exposing (onInput, onClick)
import ParserTcTurtle exposing (read, TurtleProgram)
import Drawing




-- MODEL

type alias Model =
    { input : String
    , program : TurtleProgram
    }


init : () -> ( Model, Cmd Msg )
init _ =
    ( { input = "", program = [] }
    , Cmd.none
    )


-- MESSAGES

type Msg
    = UpdateInput String
    | Parse


-- UPDATE

update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        UpdateInput str ->
            ( { model | input = str }, Cmd.none )

        Parse ->
            case read model.input of
                Ok prog ->
                    ( { model | program = prog }, Cmd.none )

                Err _ ->
                    ( { model | program = [] }, Cmd.none )


-- VIEW

view : Model -> Html Msg
view model =
    div
        [ style "font-family" "sans-serif"
        , style "padding" "20px"
        ]
        [ h2 [] [ text "TcTurtle Visualizer" ]

        , textarea
            [ value model.input
            , onInput UpdateInput
            , rows 6
            , cols 60
            , style "font-size" "16px"
            , style "margin-bottom" "10px"
            ]
            []

        , div [ style "margin" "10px 0" ]
            [ button
                [ onClick Parse
                , style "padding" "10px 20px"
                , style "font-size" "16px"
                ]
                [ text "Dessiner" ]
            ]

        , div [ style "margin-top" "20px" ]
            [ text ("Instructions parsées : " ++ Debug.toString model.program) ]

        , div
            [ style "margin-top" "20px"
            , style "display" "flex"
            , style "justify-content" "center"
            ]
            [ Drawing.display model.program ]
        ]


-- MAIN

main : Program () Model Msg
main =
    Browser.element
        { init = init
        , update = update
        , view = view
        , subscriptions = \_ -> Sub.none
        }
