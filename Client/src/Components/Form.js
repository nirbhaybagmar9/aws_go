import React, { Component } from 'react'
import {Form, Button} from 'react-bootstrap'
import '../Styles/form-style.css'
import axios from 'axios'
import {Alert} from 'react-bootstrap';

export class Forms extends Component {
    constructor(props) {
        super(props)
        this.state = {
            validated: false
        }
        this.imageId = React.createRef(); 
        this.region = React.createRef(); 
        this.name = React.createRef();
     }
    async createInstance () {
        await axios.post('http://localhost:8080/create', {
            imageId: this.imageId.current.value,
            region: this.region.current.value,
            name: [this.name.current.value],
        }).then(res => {
          console.log('Called');
        }).catch(err => console.log(err))
    }
    handleSubmit (event) {
        const form = event.currentTarget;
        event.preventDefault();

        if (form.checkValidity() === false) {
          event.stopPropagation();
          event.preventDefault();
          return
        } 
       this.createInstance();
        return false;
      };
    render() {
        const region = [
            'ap-south-1',
            'eu-west-1',
            'us-east-2',
            'ap-southeast-1',
            'ap-southeast-2',
            'eu-central-1',
            'ap-northeast-2',
            'ap-northeast-1',
            'us-east-1',
            'sa-east-1',
            'us-west-1',
            'us-west-2'
        ];
        return (
            <div className="form-container">
            <h2 className="h2">Create Instance</h2>
            <Form className="align-items-center" noValidate validated={this.state.validated}>
                <Form.Group controlId="exampleForm.ControlSelect1">
                    <Form.Label>Region</Form.Label>
                    <Form.Control ref={this.region} as="select">
                        {region.map(item => {
                            return <option key={item}>{item}</option>
                        })}
                    </Form.Control>
                </Form.Group>
              <Form.Group  md="4" controlId="validationImageId">
                <Form.Label>Image ID (AMI ID)</Form.Label>
                <Form.Control
                ref={this.imageId}
                  name="imageId"
                  required
                  type="text"
                  placeholder="Image ID"
                />
                <Form.Control.Feedback type="invalid">Enter a valid Image ID</Form.Control.Feedback>
              </Form.Group>
              <Form.Group  md="4" controlId="validationtagKey">
                <Form.Label>Instance Name</Form.Label>
                <Form.Control
                 ref={this.name}
                  required
                  type="text"
                  placeholder="Instance Name"
                />
                <Form.Control.Feedback type="invalid">Enter Instance Name</Form.Control.Feedback>
              </Form.Group>
            <Button type="button" onClick={(e) => this.handleSubmit(e)}>Create Instance</Button>
          </Form>
          </div>
        );
      }
}

