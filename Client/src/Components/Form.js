import React, { Component } from 'react'
import {Form, Button} from 'react-bootstrap'
import '../Styles/form-style.css'
import axios from 'axios'
export class Forms extends Component {
    constructor(props) {
        super(props)
        this.state = {
            validated: false
        }
        this.imageId = React.createRef(); 
        this.tagKey = React.createRef(); 
        this.tagValue = React.createRef(); 
        this.region = React.createRef(); 

     }
    async createInstance () {
        await axios.post('http://localhost:8080/create', {
            imageId: this.imageId.current.value,
            region: this.region.current.value,
            tagKey: this.tagKey.current.value,
            tagValue: this.tagValue.current.value
        }).then(res => console.log(res)).catch(err => console.log(err))
    }
    handleSubmit (event) {
        const value = this.imageId.current.value
        console.log(value)
        const form = event.currentTarget;
        console.log(event.target)
        if (form.checkValidity() === false) {
          event.preventDefault();
          event.stopPropagation();
        }
        this.setState({
            validated: true
        }, () => this.createInstance())
        event.stopPropagation();
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
            <Form className="align-items-center" noValidate validated={this.state.validated} onSubmit={this.handleSubmit.bind(this)}>
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
                <Form.Label>Tag Key</Form.Label>
                <Form.Control
                 ref={this.tagKey}
                  required
                  type="text"
                  placeholder="Tag Key"
                />
                <Form.Control.Feedback type="invalid">Enter Tag Key</Form.Control.Feedback>
              </Form.Group>
              <Form.Group  md="4" controlId="validationTagValue">
                <Form.Label>Tag Value</Form.Label>
                  <Form.Control
                    ref={this.tagValue}
                    type="text"
                    placeholder="Tag Value"
                    required
                  />
                  <Form.Control.Feedback type="invalid">Enter Tag Value</Form.Control.Feedback>
              </Form.Group>
            <Button type="submit">Create Instance</Button>
          </Form>
          </div>
        );
      }
}

