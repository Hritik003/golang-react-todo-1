import React,{Component} from "react";
import axios from 'axios';
import {Card, Header, Form, Input, Icon } from 'semantic-ui-react';

let endpoint = "http://localhost:9000";

class TodoList extends Component{
    constructor(props){
        super(props);

        this.state={
            task:"",
            items:[],
        }
    }

    componentDidMount(){
        this.getTask();
    }

    render(){
        return (
            <div>
                <div className="row">

                </div>
            </div>
        );
    }
} 

export default TodoList;



 