import React from 'react';
import axios from 'axios'
import {Modal, Button } from 'react-bootstrap';
import { Form } from 'tabler-react'

export class CustomModal extends React.Component {
  constructor(props) {
    super(props);
    this.handleClose = this.handleClose.bind(this);
    this.newTransaction = this.newTransaction.bind(this);
    this.handleShow = this.handleShow.bind(this);
  }

  static getDerivedStateFromProps(props, state) {
    if (props.object.show !== state.showModal) {
      return {
        showModal: props.object.show,
        senderPublicKey:props.object.senderPublicKey,
        senderPrivateKey:props.object.senderPrivateKey,
        recipientPublicKey:props.object.recipientPublicKey,
        signature:props.object.signature,
        amount:props.object.amount,
      };
    }
    // Return null to indicate no change to state.
    return null;
  }

  state ={
    "showModal" : false,
    "senderPublicKey":null,
    "senderPrivateKey":null,
    "recipientPublicKey":null,
    "amount":null,
    "signature":null,
    "blockchainNode":"http://127.0.0.1:5001"
  }
  handleClose() {
    this.props.action(false)
  }

  handleShow() {
    this.setState({showModal: true})
  }

  newTransaction() {
    axios.defaults.headers['Access-Control-Allow-Methods'] = 'GET, POST';
    axios.defaults.headers['Access-Control-Allow-Origin'] = '*';
    axios.post(this.state.blockchainNode + '/transaction/new', 
    {
      "senderPublicKey":this.state.senderPublicKey,
      "senderPrivateKey":this.state.senderPrivateKey,
      "recipientPublicKey":this.state.recipientPublicKey,
      "amount":this.state.amount,
    }, 
      )
      .then((response) => {
        console.log(response.data);
        this.setState({blockchainNode: response.data.data})
        // this.setState({ show: true, signature: response.data.Signature}, () => {console.log(this.state)});
        // this.setState({
        //   privateKey: response.data.privateKey,
        //   publicKey: response.data.publicKey
        // }, function () {
        //   console.log(this.state.privateKey);
        // });
      }
      )
  }

  render (){
    return (
      <>
      <Modal show={this.state.showModal} onHide={this.handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>Confirm transaction</Modal.Title>
        </Modal.Header>
        <Modal.Body>
        <p>
        Sender Public Key:
        </p>
    <Form.Input name='senderPublicKey' value={this.state.senderPublicKey} readOnly={true}/>
    <p>
        Recipient Public Key:
        </p>
    <Form.Input name='recipientPublicKey' value={this.state.recipientPublicKey} readOnly={true}/>
    <p>
        Amount:
        </p>
    <Form.Input name='amount' value={this.state.amount} readOnly={true}/>
    <p>
        Signature:
        </p>
    <Form.Input name='signature' value={this.state.signature} readOnly={true}/>
    <p>
        Blockchain Node:
        </p>
    <Form.Input name='blockchainNode' value={this.state.blockchainNode}/>

        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={this.handleClose}>
            Close
          </Button>
          <Button variant="success" onClick={this.newTransaction}>
            Confirm
          </Button>
        </Modal.Footer>
      </Modal>
    </>
    );
  }
}