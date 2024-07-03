# Invoicetastic
Invoicetastic is a CLI-based tool for generating PDF invoices.

# Installation

In order to use Invoicetastic, first you need to install `texlive` tool in your system

Then, build it from sources:

```bash
$ cd invoicestastic/
$ go build
```

That will generate an executable `invoicetastic`.

# Usage

## Create empty invoice YAML file

It creates a `.yml` file in the current directory with an empty invoice.

```bash
$ invoicetastic create [invoice number]
```

**Example**
```bash
$ invoicetastic create I-2024-3
I-2024-3.yml file created
```

### Default invoice data
Some information is common to all invoices you create. To avoid re-entering this information each time, create a configuration file named `invoicetastic.yml` in your current directory. In this file, add a `defaults` section to store default invoice data (e.g., provider, contact, and other necessary details).


When you run `invoicetastic create`, invoice data from `invoicetastic.yml` will be automatically included in the created invoice.

[Example of `invoicetastic.yml` file](./examples/invoicetastic.yml)

## Generate PDF from invoice YAML file

It renders invoice data to PDF using the given template (see section Templates). Also, it computes product and invoice totals, taking taxes into account. 

```bash
$ invoicetastic generate -f [invoice YAML file or URL] -t [template path or URL]
```
**Example**
```bash
$ invoicetastic generate -f I-2024-3.yml -t https://raw.githubusercontent.com/hvalls/invoicetastic/main/_templates/english-usd.tex
I-2024-3.pdf file created
```

#### Parameters 
| Parameter | Description | Required | Default value |
|---|---|---|---|
| `-f, --file`   | Invoice YAML file path or URL  |  Yes |  |
| `-t, --template`   | Invoice template (.tex) file path or URL   |  No | https://raw.githubusercontent.com/hvalls/invoicetastic/main/_templates/english-usd.tex |

# Invoice YAML file

[This is an example of an invoice YAML file](./examples/invoice.example.yml)

### Import entities

You can import `provider`, `customer`, `taxes`, `contact` and `paymentInfo` sections from other YAML files so you don't have to rewrite them in every invoice. 

For example, create a file `./customers/acme.yml` with content below:

```yaml
name: ACME Inc.
vat:  123456789
address: 
  line1: This is
  line2: the 
  line3: address
```

and then import it inside your invoice main YAML file:

```yaml
// ...
customer:
  import: "./customers/acme.yml"
// ...
```

NOTE: Remember `taxes` YAML file must be an array.

# Templates

A template is a LaTex file using [Golang template system](https://pkg.go.dev/text/template). There are some templates under [_templates](https://github.com/hvalls/invoicetastic/tree/main/_templates) directory you can use. Also, you can create your own. 

# Parameters mapping

| Template file (.tex) parameter | Invoice YAML file parameter  | Description  | 
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
| `{{ .Products[*].Total }}`   | Computed value  | Product total |
| `{{ .Subtotal }}`   | Computed value | Invoice subtotal |
| `{{ .Taxes }}`   | `taxes`  |  List of taxes |
| `{{ .Taxes[*].Name }}`   | `taxes.*.name`  |  Tax name |
| `{{ .Taxes[*].Percentage }}`   | `taxes.*.percentage`  |  Tax percentage |
| `{{ .Taxes[*].Value }}`   | Computed value  |  Tax value |
| `{{ .Total }}`   | Computed value  |  Invoice total |
| `{{ .Contact.Name }}`   | `contact.name`  |  Contact name |
| `{{ .Contact.Email }}`   | `contact.email`  |  Contact email |
| `{{ .Contact.Website }}`   | `contact.website`  |  Contact website |
| `{{ .PaymentInfo.Bank }}`   | `paymentInfo.bank`  |  Bank name |
| `{{ .PaymentInfo.AccountName }}`   | `paymentInfo.accountName`  |  Account name |
| `{{ .PaymentInfo.AccountNumber }}`   | `paymentInfo.accountNumber`  |  Account number |
| `{{ .PaymentInfo.SwiftBIC }}`   | `paymentInfo.swiftBic`  |  SWIFT/BIC |
