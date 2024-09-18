# Invoicetastic
Invoicetastic is a CLI-based tool for generating PDF invoices from LaTex templates.

## Installation

In order to use Invoicetastic, first you need to install `texlive` tool in your system.

Then, download the binary for your operating system from the [Releases page](https://github.com/hvalls/invoicetastic/releases).

## Usage

### invoicetastic create-provider

It creates a [provider YAML file](examples/provider.yml) in the current directory.

**Parameters**
| Parameter | Description | Required |
|---|---|---|
| `--name`   | Provider name  |  Yes |
| `--vatnum`   | Provider VAT number  |  No |
| `--address`   |  Address line. Multiple values are allowed.  |  No |

**Example**
```bash
$ invoicetastic create-provider \
    --name "Tech Innovators Inc." \
    --vatnum GB123456789 \
    --address "1234 Innovation Drive, Suite 567" \
    --address "Silicon Valley, CA 94043"
```

### invoicetastic create-client

It creates a [client YAML file](examples/client.yml) in the current directory.

**Parameters**
| Parameter | Description | Required |
|---|---|---|
| `--name`   | Client name  |  Yes |
| `--vatnum`   | Client VAT number  |  No |
| `--address`   |  Address line. Multiple values are allowed.  |  No |

**Example**
```bash
$ invoicetastic create-client \
    --name "Creative Solutions Ltd." \
    --vatnum DE987654321 \
    --address "789 Pioneer Avenue, Floor 3" \
    --address "New York"
```

### invoicetastic create-contact

It creates a [contact YAML file](examples/contact.yml) in the current directory.

**Parameters**
| Parameter | Description | Required |
|---|---|---|
| `--name`   | Contact name  |  Yes |
| `--email`   | Contact email address  |  No |
| `--website`   |  Contact website  |  No |

**Example**
```bash
$ invoicetastic create-contact \
    --name "Creative Solutions Ltd." \
    --email "jhon.doe@mail.com" \
    --website "johndoe.com"
```

### invoicetastic create-product

It creates a [product YAML file](examples/product1.yml) in the current directory.

**Parameters**
| Parameter | Description | Required |
|---|---|---|
| `--name`   | Product name  |  Yes |
| `--unitprice`   | Product unit price  |  No |

**Example**
```bash
$ invoicetastic create-product \
    --name "Innovative Tech Software License" \
    --unitprice 500
```


### invoicetastic create-tax

It creates a [tax YAML file](examples/tax.yml) in the current directory.

**Parameters**
| Parameter | Description | Required |
|---|---|---|
| `--name`   | Tax name  |  Yes |
| `--percentage`   | Tax percentage  |  No |

**Example**
```bash
$ invoicetastic create-tax \
    --name "VAT" \
    --percentage 10
```

### invoicetastic create-paymentinfo

It creates a [payment info YAML file](examples/paymentInfo.yml) in the current directory.

**Parameters**
| Parameter | Description | Required |
|---|---|---|
| `--bank`   | Bank name  |  Yes |
| `--accountname`   | Bank account name  |  No |
| `--accountnum`   |  Bank account number  |  No |
| `--swiftbic`   |  Bank Swift/BIC  |  No |

**Example**
```bash
$ invoicetastic create-paymentinfo \
    --bank "Silicon Valley Bank" \
    --accountname "Tech Innovators Inc." \
    --accountnum "123456789" \
    --swiftbic "SVBXXX"
```

### invoicetastic create-invoice [invoice number]

It creates a [invoice YAML file](examples/invoice.yml) in the current directory.

**Parameters**
| Parameter | Description | Required |
|---|---|---|
| `[invoice number]`   | Invoice number/code  |  Yes |
| `--date`   | Invoice date  |  No |
| `--duedate`   | Invoice due date  |  No |
| `--provider`   | Provider YAML file path or URL  |  No |
| `--client`   | Client YAML file path or URL  |  No |
| `--line`   |  Invoice line, composed of a product YAML file path and a quantity, separated by colon. e.g. `product1.yml:2`. Multiple values are allowed.  |  No |
| `--tax`   | Tax YAML file path or URL. Multiple values are allowed.  |  No |
| `--contact`   | Contact YAML file path or URL  |  No |
| `--payment`   | PaymentInfo YAML file path or URL  |  No |

**Example**
```bash
$ invoicetastic create-invoice I-2024-3 \ 
    --date "2024-09-01" \
    --duedate "2024-10-01" \
    --provider "examples/provider.yml" \
    --client "examples/client.yml" \
    --line "examples/product1.yml:2" \
    --line "examples/product2.yml:4" \
    --tax "examples/tax.yml" \
    --contact "examples/contact.yml" \
    --payment "examples/paymentInfo.yml"
```

### invoicetastic render

It renders invoice data to PDF using the template provided. (see section Templates). Also, it computes product and invoice totals, taking taxes into account. 

**Parameters**
| Parameter | Description | Required |
|---|---|---|
| `--file`   | Invoice YAML file path or URL  |  Yes |  |
| `--template`   | Invoice template (.tex) file path or URL  |  Yes | 

**Example**
```bash
$ invoicetastic render --file I-2024-3.yml --template https://raw.githubusercontent.com/hvalls/invoicetastic/main/_templates/english-usd.tex
```

### Templates

A template is a LaTex file using [Golang template system](https://pkg.go.dev/text/template). There are a couple of templates under [_templates](https://github.com/hvalls/invoicetastic/tree/main/_templates) directory you can use. Also, you can create your own. 

#### Parameters mapping

| Template file (.tex) parameter | Invoice YAML file parameter  | Description  | 
|---|---|---|
| `{{ .Number }}`   | `number`  |  Invoice number |
| `{{ .Date }}`   | `date`  |  Invoice date |
| `{{ .DueDate }}`   | `dueDate`  |  Invoice due date |
| `{{ .Provider.Name }}`   | `provider.name`  |  Provider name |
| `{{ .Provider.VATNumber }}`   | `provider.vatNumber`  |  Provider VAT number |
| `{{ .Provider.Address }}`   | `provider.address`  |  Provider address lines |
| `{{ .Client.Name }}`   | `client.name`  |  Client name |
| `{{ .Client.VATNumber }}`   | `client.vatNumber`  |  Client VAT number |
| `{{ .Client.Address }}`   | `client.address`  |  Client address lines |
| `{{ .Lines }}`   | `lines`  |  Invoice lines |
| `{{ .Lines[*].Name }}`   | `lines.*.name`  | Line's product name |
| `{{ .Lines[*].Qty }}`   | `lines.*.qty`  | Line's product quantity |
| `{{ .Lines[*].UnitPrice }}`   | `lines.*.unitPrice`  | Line's product unit price |
| `{{ .Lines[*].Total }}`   | Computed value  | Line total |
| `{{ .Subtotal }}`   | Computed value | Invoice subtotal |
| `{{ .Taxes }}`   | `taxes`  |  Invoice taxes list |
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
