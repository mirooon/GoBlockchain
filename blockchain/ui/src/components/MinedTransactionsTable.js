import React from 'react';
import { Table } from 'tabler-react'

export class MinedTransactionsTable extends React.Component {
static getDerivedStateFromProps(props, state) {
    if (props.transactions !== state.transactions) {
      return {
        transactions: props.transactions,
      };
    }
    return null;
  }
  state = {
    "transactions": this.props.transactions,
  }
    render() {
    return (<Table>
        <Table.Header>
          <Table.ColHeader>#</Table.ColHeader>
          <Table.ColHeader>Sender Public Key</Table.ColHeader>
          <Table.ColHeader>Recipient Public Key</Table.ColHeader>
          <Table.ColHeader>Amount</Table.ColHeader>
          <Table.ColHeader>Timestamp</Table.ColHeader>
          <Table.ColHeader>Block Number</Table.ColHeader>
        </Table.Header>
        <Table.Body>
        {this.state.transactions != null && this.state.transactions.length > 0 ? (
            <>
            {this.state.transactions.map(function(transaction, index){
              
                          return (<><Table.Row>
                              <Table.Col>{index+1}</Table.Col>
                          <Table.Col>{transaction.SenderPublicKey.substring(0, 10)}...{transaction.SenderPublicKey.substring(transaction.SenderPublicKey.length-11, transaction.SenderPublicKey.length-1)}</Table.Col>
                          <Table.Col>{transaction.RecipientPublicKey.substring(0, 10)}...{transaction.RecipientPublicKey.substring(transaction.RecipientPublicKey.length-11, transaction.RecipientPublicKey.length-1)}</Table.Col>
                          <Table.Col>{transaction.Amount}</Table.Col>
                          <Table.Col>{transaction.Timestamp}</Table.Col>
                          <Table.Col>{transaction.BlockNumber}</Table.Col>
                          </Table.Row></>
                        )}
                        )
                    }
                    </>
                       ) : (
                         <><Table.Col>No transactions</Table.Col></>
                       )}
                    
                    </Table.Body>
                  </Table>);
    }
};