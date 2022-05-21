import React, {ReactNode} from "react";
import {TemplateExample} from "../templates/templateExample";
import "../styles/styleExample.scss"

interface ExampleState {
    example: string;
}

interface ExampleProps {
}

export interface ExampleInterface {
    state: ExampleState;
    props: ExampleProps

    exampleFunc(): void;
}

export class Example extends React.Component<ExampleProps, ExampleState> implements ExampleInterface{
    state = {
        example: "example template"
    }

    exampleFunc() {
        console.log('success')
    }

    render(): ReactNode {
        return TemplateExample(this)
    }
}