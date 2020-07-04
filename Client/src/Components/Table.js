import React from 'react'
import '../Styles/table-style.css'
import {Dropdown} from 'react-bootstrap';
import axios from 'axios';
import { CheckboxToggle } from 'react-rainbow-components';


export class Table extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
           instances: [],
           region: 'us-east-2',
        }
     }
    async getInstances() {
        await axios.get(`http://localhost:8080/?region=${this.state.region}`)
        .then(res => {
        const instances = res.data;
        console.log(res)
        this.setState({ instances : instances });   
        }).catch(err => console.log(err))
    }
    async componentDidMount () {
      await this.getInstances();
    }

    selectRegion (key) {        
        this.setState({
            region: key
        }, () => this.getInstances() )
    }

    async startStopInstance (val, instance) {
        let state = false
        if (val === false){
            state = 'START'
        } else {
            state = 'STOP'
        }
        await axios.post(`http://localhost:8080/state`, {
             region: this.state.region,
             instanceID: instance.InstanceID,
             state: state
        })
        .then(res => {
            console.log(res);
            this.getInstances()
            setTimeout(this.getInstances.bind(this), 60000)
        }).catch(err => console.log(err))
    }

    render() {
        const region = [
            'us-east-2',
            'ap-south-1',
            'eu-west-1',
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
          <div className="table-container">
              <Dropdown>
                <Dropdown.Toggle variant="success" id="dropdown-basic">
                    Region
                </Dropdown.Toggle>
                <Dropdown.Menu>
                    {region.map(item => {
                        return  <Dropdown.Item key={item} eventKey={item} onSelect={eventKey => this.selectRegion(eventKey)}>{item}</Dropdown.Item>
                    })}
                </Dropdown.Menu>
                </Dropdown>
             <h1 id='title'>AWS EC2 Instances</h1>
             <table id='instances'>
                <tbody>
                   <tr>
                       <th>Instance ID</th>
                       <th>Image ID</th>
                       <th>Instance Type</th>
                       <th>Availabilty Status</th>
                       <th>Status</th>
                   </tr>
                   {this.state.instances.map((instance, index) => {
                       let value = instance.State === 'stopped' || instance.State === 'stopping' || instance.State === 'terminated' ? false : true;
                       let disabled = instance.State === 'terminated' ? true : false
                        return (
                            <tr key={instance.InstanceID}>
                                <td>{instance.InstanceID}</td>
                                <td>{instance.ImageID}</td>
                                <td>{instance.InstanceType}</td>
                                <td>{instance.AvailabilityZone}</td>
                                <td>
                                    <CheckboxToggle
                                        disabled={disabled}
                                        label={instance.State}
                                        value={value}
                                        onChange={() => this.startStopInstance(value, instance)}
                                    />
                                </td>
                            </tr>
                        );
                })}
                </tbody>
             </table>
          </div>
       )
    }
 }