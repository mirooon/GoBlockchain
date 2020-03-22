import React from 'react';
import axios from 'axios'
import { Grid, Header, Button, Alert, Form } from 'tabler-react'

export class Neighbours extends React.Component {
  constructor(props) {
    super(props);
    this.handleChange = this.handleChange.bind(this);
    this.getNeighbours = this.getNeighbours.bind(this);
    this.registerNeighbour = this.registerNeighbour.bind(this);
  }

  state = {
    "neighbours": [],
    "node": null,
    "message": null
  }
  
  componentDidMount(){
    this.getNeighbours();
  }

  handleChange (evt) {
    this.setState({ [evt.target.name]: evt.target.value });
  }

  getNeighbours() {
    axios.get('http://localhost:5001/nodes')
      .then((response) => {
          this.setState({ neighbours: response.data.Nodes}, () => {console.log("neigbours downloaded")});
          console.log(response.data.Nodes);
      }
      ).catch((err) => {
        this.setState({ errorMessage: "Problem with download the neighbours!", message: null});
        setTimeout( () => {
            this.setState( () => ({
                errorMessage: null
            }));
          }, 5000);
      })
  }

  registerNeighbour() {
    axios.post('http://localhost:5001/node/new', {Node: this.state.node})
      .then((response) => {
          console.log('response')
          console.log(response)
        this.setState({ neighbours: response.data.AllFollowingNodes, message: response.data.Message}, () => {console.log("neigbours downloaded")});
      }
      ).catch((err) => {
        this.setState({ errorMessage: "Problem with neighbour registration!", message: null});
        setTimeout( () => {
            this.setState( () => ({
                errorMessage: null
            }));
          }, 5000);
      })
  }

  ValidateIPaddress(ipaddress) {
    if (/^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?):[0-9]{1,5}$/.test(ipaddress))
     {
       return (true)
     }
   return (false)
   }

    render() {
    return (
<div className="container-fluid">
        <Header.H1>Neighbours</Header.H1>
        {!this.ValidateIPaddress(this.state.node) &&         <p><Alert type="primary">
  <strong>Provide a node (neighbour) address with port!</strong> (example: 127.0.0.1:5001)
</Alert></p>}
        <Grid.Row cards deck>
          <Grid.Col md={4}>
          </Grid.Col>
          <Grid.Col md={4}>
          <Form.Group>
  <Form.Input name="node" placeholder="New node (neighbour) address" onChange={this.handleChange}/>
</Form.Group>
          </Grid.Col>
          <Grid.Col md={4}>
          </Grid.Col>
        </Grid.Row>
        <p><Button color="primary" onClick={this.registerNeighbour}>
    Register node
  </Button></p>
  {this.state.Message != null &&         <p><Alert type="primary">
    <strong>{this.state.Message}</strong>
</Alert></p>}
{this.state.errorMessage != null &&         <p><Alert type="danger">
    <strong>{this.state.errorMessage}</strong>
</Alert></p>}
<Grid.Row cards deck>
          <Grid.Col md={4}>
          </Grid.Col>
          <Grid.Col md={4}>
              {this.state.neighbours != null &&       <ul>
          {this.state.neighbours.map((val, idx) => {     
           return (<li>{val}</li>) 
        })}
        </ul>}
          </Grid.Col>
          <Grid.Col md={4}>
          </Grid.Col>
        </Grid.Row>
                  </div>
                  )
}
};