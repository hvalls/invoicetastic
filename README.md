# Invoicetastic
Invoicetastic is a CLI-based tool for generating invoices as PDF from a YAML file.

## Installation

In order to use Invoicetastic, first you need to install `texlive` tool in your system. 

Then, build it from sources:

```
$ cd invoicestastic/
$ go build
```

That will generate an executable `invoicetastic`.

## Usage

```bash
$ invoicetastic generate -f {path_to_invoice_yaml_file} -t {path_to_template}
```

#### Example

```bash
$ invoicetastic generate -f invoice.example.yml -t _templates/english-usd.tex
```

## Invoice YAML file format

[This is an example of an invoice YAML file](./invoice.example.yml)

## Templates

A template is just a LaTex file using [Golang template system](https://pkg.go.dev/text/template). There are some templates under `_templates` directory you can use. Also, you can create your own templates. 

### Parameters mapping

| Template file parameter | YAML file parameter  | Description  | 
|---|---|---|
| `{{ .Number }}`   | `number`  |  Invoice number |
| `{{ .Date }}`   | `date`  |  Invoice date |
| `{{ .DueDate }}`   | `dueDate`  |  Invoice due date |
| `{{ .Provider.Name }}`   | `provider.name`  |  Provider name |
| `{{ .Provider.VAT }}`   | `provider.vat`  |  Provider VAT |
| `{{ .Provider.Address.Line1 }}`   | `provider.address.line1`  |  Provider address line 1 |
| `{{ .Provider.Address.Line2 }}`   | `provider.address.line1`  |  Provider address line 2 |
| `{{ .Provider.Address.Line3 }}`   | `provider.address.line1`  |  Provider address line 3 |
| `{{ .Customer.Name }}`   | `customer.name`  |  Customer name |
| `{{ .Customer.VAT }}`   | `customer.vat`  |  Customer VAT |
| `{{ .Customer.Address.Line1 }}`   | `customer.address.line1`  |  Customer address line 1 |
| `{{ .Customer.Address.Line2 }}`   | `customer.address.line1`  |  Customer address line 2 |
| `{{ .Customer.Address.Line3 }}`   | `customer.address.line1`  |  Customer address line 3 |
| `{{ .Products }}`   | `products`  |  List of products |
| `{{ .Products[*].Description }}`   | `products.*.description`  | Product description |
| `{{ .Products[*].Qty }}`   | `products.*.qty`  | Product quantity |
| `{{ .Products[*].UnitPrice }}`   | `products.*.unitPrice`  | Product unit price |
| `{{ .Products[*].Total }}`   | `products.*.total`  | Product total |
| `{{ .Subtotal }}`   | `subtotal`  | Invoice subtotal |
| `{{ .Taxes }}`   | `taxes`  |  List of taxes |
| `{{ .Taxes[*].Name }}`   | `taxes.*.name`  |  Tax name |
| `{{ .Taxes[*].Percentage }}`   | `taxes.*.percentage`  |  Tax percentage |
| `{{ .Taxes[*].Value }}`   | `taxes.*.value`  |  Tax value |
| `{{ .Total }}`   | `total`  |  Invoice total |
| `{{ .Contact.Name }}`   | `contact.name`  |  Contact name |
| `{{ .Contact.Email }}`   | `contact.email`  |  Contact email |
| `{{ .Contact.Website }}`   | `contact.website`  |  Contact website |
| `{{ .PaymentInfo.Bank }}`   | `paymentInfo.bank`  |  Bank name |
| `{{ .PaymentInfo.AccountName }}`   | `paymentInfo.accountName`  |  Account name |
| `{{ .PaymentInfo.AccountNumber }}`   | `paymentInfo.accountNumber`  |  Account number |
| `{{ .PaymentInfo.SwiftBIC }}`   | `paymentInfo.swiftBic`  |  SWIFT/BIC |
