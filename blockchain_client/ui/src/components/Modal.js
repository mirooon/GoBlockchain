import React from 'react';
import axios from 'axios'
import {Modal, Button } from 'react-bootstrap';
import { Form } from 'tabler-react'

export class CustomModal extends React.Component {
  constructor(props) {
    super(props);
    this.handleChange = this.handleChange.bind(this);
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
    "blockchainNode":null,
    "verifyResult":null
  }

  handleChange (evt) {
    this.setState({ [evt.target.name]: evt.target.value });
  }
  
  handleClose() {
    this.props.action(false)
    this.setState({verifyResult: null})
  }

  handleShow() {
    this.setState({showModal: true})
  }

  newTransaction() {
    axios.defaults.headers['Access-Control-Allow-Methods'] = 'GET, POST';
    axios.defaults.headers['Access-Control-Allow-Origin'] = '*';
    const body = {
      "senderPublicKey":this.state.senderPublicKey,
      "recipientPublicKey":this.state.recipientPublicKey,
      "signature":this.state.signature,
      "amount":parseFloat(this.state.amount),
      "blockchainNode": this.state.blockchainNode
    };
    axios.post('http://localhost:8080/transaction/submit', 
    body, 
      )
      .then((response) => {
        this.setState({ verifyResult: response.data.verifyResult});
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
        {this.state.verifyResult == null ? (
          <>
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
            <Form.Input name='blockchainNode' placeholder="ex. 172.16.1.12:5001" onChange={this.handleChange}/>
                </Modal.Body>
                <Modal.Footer>
                  <Button variant="secondary" onClick={this.handleClose}>
                    Close
                  </Button>
                  <Button variant="success" onClick={this.newTransaction}>
                    Confirm
                  </Button>
                </Modal.Footer>
                </>
      ) : (
        this.state.verifyResult === "true" ? (
          <>
                <Modal.Body>
                <p>
                Success
                </p>
                </Modal.Body>
                <Modal.Footer>
                  <Button variant="secondary" onClick={this.handleClose}>
                    Close
                  </Button>
                </Modal.Footer>
                </>
      ) : (
        <>
        <Modal.Body>
        <p>
        Failed
        </p>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={this.handleClose}>
            Close
          </Button>
        </Modal.Footer>
        </>
      )
      )}
      </Modal>
    </>
    );
  }
}