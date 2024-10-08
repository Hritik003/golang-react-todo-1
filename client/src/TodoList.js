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
                    <Header className="header" as="h2" color="yellow">
                        My To Do List!
                    </Header>   
                </div>
            </div>
        );
    }
} 

export default TodoList;



 