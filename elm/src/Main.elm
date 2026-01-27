module Main exposing (main)

import Browser
import Html exposing (Html, div, text, textarea, button, h2, h3, h4, p, span, ul, li)
import Html.Attributes exposing (value, style, rows, cols, class)
import Html.Events exposing (onInput, onClick)
import ParserTcTurtle exposing (read, TurtleProgram)
import Drawing


-- MODEL

type alias Model =
    { input : String
    , program : TurtleProgram
    , showHelp : Bool
    , error : Maybe String
    }


init : () -> ( Model, Cmd Msg )
init _ =
    ( { input = "", program = [], showHelp = True, error = Nothing }
    , Cmd.none
    )



-- MESSAGES

type Msg
    = UpdateInput String
    | Parse
    | ToggleHelp


-- UPDATE

update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        UpdateInput str ->
            ( { model | input = str }, Cmd.none )

        Parse ->
            case read model.input of
                Ok prog ->
                    ( { model | program = prog, showHelp = False, error = Nothing }, Cmd.none )

                Err err ->
                    ( { model | program = [], error = Just "Erreur de syntaxe ! Vérifiez votre code." }, Cmd.none )

        ToggleHelp ->
            ( { model | showHelp = not model.showHelp }, Cmd.none )


-- VIEW

view : Model -> Html Msg
view model =
    div
        [ style "font-family" "Arial, sans-serif"
        , style "max-width" "1200px"
        , style "margin" "0 auto"
        , style "padding" "20px"
        , style "background-color" "#f5f5f5"
        ]
        [ header
        , div [ style "display" "flex", style "gap" "20px" ]
            [ leftPanel model
            , rightPanel model
            ]
        ]


header : Html Msg
header =
    div
        [ style "margin-bottom" "30px"
        , style "padding-bottom" "20px"
        , style "border-bottom" "3px solid #3498db"
        ]
        [ h2
            [ style "color" "#2c3e50"
            , style "margin" "0 0 10px 0"
            ]
            [ text "🐢 TcTurtle Visualizer" ]
        ]


leftPanel : Model -> Html Msg
leftPanel model =
    div
        [ style "flex" "1"
        , style "display" "flex"
        , style "flex-direction" "column"
        , style "gap" "15px"
        ]
        [ editorSection model
        , helpSection model
        ]


editorSection : Model -> Html Msg
editorSection model =
    div
        [ style "background-color" "white"
        , style "padding" "20px"
        , style "border-radius" "8px"
        , style "box-shadow" "0 2px 4px rgba(0,0,0,0.1)"
        ]
        [ h3
            [ style "margin-top" "0"
            , style "color" "#2c3e50"
            ]
            [ text "📝 Éditeur" ]
        , textarea
            [ value model.input
            , onInput UpdateInput
            , rows 10
            , cols 40
            , style "width" "100%"
            , style "padding" "10px"
            , style "font-family" "monospace"
            , style "font-size" "14px"
            , style "border" "1px solid #bdc3c7"
            , style "border-radius" "4px"
            , style "box-sizing" "border-box"
            ]
            []
        , if model.error /= Nothing then
            div
                [ style "margin-top" "15px"
                , style "padding" "12px"
                , style "background-color" "#fee"
                , style "border-left" "4px solid #e74c3c"
                , style "border-radius" "4px"
                , style "color" "#c0392b"
                , style "font-weight" "bold"
                ]
                [ text "❌ ", text (Maybe.withDefault "" model.error) ]
          else
            text ""
        , div [ style "margin-top" "15px" ]
            [ button
                [ onClick Parse
                , style "background-color" "#3498db"
                , style "color" "white"
                , style "padding" "12px 24px"
                , style "font-size" "16px"
                , style "border" "none"
                , style "border-radius" "4px"
                , style "cursor" "pointer"
                , style "font-weight" "bold"
                ]
                [ text "▶ Dessiner" ]
            ]
        ]


helpSection : Model -> Html Msg
helpSection model =
    div
        [ style "background-color" "white"
        , style "padding" "20px"
        , style "border-radius" "8px"
        , style "box-shadow" "0 2px 4px rgba(0,0,0,0.1)"
        ]
        [ button
            [ onClick ToggleHelp
            , style "background-color" "#27ae60"
            , style "color" "white"
            , style "padding" "10px 20px"
            , style "border" "none"
            , style "border-radius" "4px"
            , style "cursor" "pointer"
            , style "font-weight" "bold"
            , style "width" "100%"
            ]
            [ text (if model.showHelp then "▼ Masquer l'aide" else "▶ Afficher l'aide") ]
        , if model.showHelp then
            div [ style "margin-top" "15px" ]
                [ h4 [ style "margin-top" "0" ] [ text "Syntaxe du langage :" ]
                , ul [ style "padding-left" "20px" ]
                    [ li [] [ span [ style "font-weight" "bold" ] [ text "Forward N" ], text " : Avancer de N pixels" ]
                    , li [] [ span [ style "font-weight" "bold" ] [ text "Left N" ], text " : Tourner à gauche de N degrés" ]
                    , li [] [ span [ style "font-weight" "bold" ] [ text "Right N" ], text " : Tourner à droite de N degrés" ]
                    , li [] [ span [ style "font-weight" "bold" ] [ text "Repeat N [...]" ], text " : Répéter N fois les instructions" ]
                    ]
                , h4 [] [ text "Format :" ]
                , p [] [ text "Les instructions doivent être dans des crochets [ ], séparées par des virgules" ]
                , h4 [] [ text "Exemple :" ]
                , p
                    [ style "background-color" "#ecf0f1"
                    , style "padding" "10px"
                    , style "border-radius" "4px"
                    , style "font-family" "monospace"
                    , style "font-size" "12px"
                    ]
                    [ text "[Forward 100, Right 90, Forward 100, Right 90, Forward 100, Right 90, Forward 100]" ]
                ]
          else
            text ""
        ]


rightPanel : Model -> Html Msg
rightPanel model =
    div
        [ style "flex" "1"
        , style "background-color" "white"
        , style "padding" "20px"
        , style "border-radius" "8px"
        , style "box-shadow" "0 2px 4px rgba(0,0,0,0.1)"
        , style "display" "flex"
        , style "flex-direction" "column"
        , style "align-items" "center"
        ]
        [ h3
            [ style "margin-top" "0"
            , style "color" "#2c3e50"
            , style "width" "100%"
            ]
            [ text "🎨 Résultat" ]
        , div
            [ style "flex" "1"
            , style "display" "flex"
            , style "align-items" "center"
            , style "justify-content" "center"
            , style "width" "100%"
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
