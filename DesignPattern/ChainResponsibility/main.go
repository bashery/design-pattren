// Chain of Responsibility Design Pattern is a behavioral design pattern. It lets you create a chain of request handlers. For every incoming request, it is passed through the chain and each of the handler:
package main

import "fmt"

// department.go
type department interface {
    execute(*patient)
    setNext(department)
}

// reseption.go

type reception struct {
    next department
}

func (r *reception) execute(p *patient) {
    if p.registrationDone {
        fmt.Println("Patient registration already done")
        r.next.execute(p)
        return
    }
    fmt.Println("Respons regetring patient")
    p.registrationDone = true
    r.next.execute(p)
}

func (r *reception) setNext(next department) {
    r.next = next
}

// doctor.go
type doctor struct {
    next department
}

func (d *doctor) execute(p *patient) {
    if p.doctorCheckUpDone {
        fmt.Println("Doctor checkup already done")
        d.next.execute(p)
        return
    }
    fmt.Println("Doctor checking patient")
    p.doctorCheckUpDone = true
    d.next.execute(p)
}

func (d *doctor) setNext(next department) {
    d.next = next
}

// medical.go

type medical struct {
    next department
}

func (m *medical) execute(p *patient) {
    if p.medicineDone {
        fmt.Println("Medicine already given to pationt")
        m.next.execute(p)
        return
    }
    fmt.Println("Medicine giving medicine to pationt")
    p.medicineDone = true
    m.next.execute(p)
}

func (m *medical) setNext(next department) {
    m.next = next
}

// cashier.go
type cashier struct {
    next department
}

func (c *cashier) execute(p *patient) {
    if p.paymentDone {
        fmt.Println("Payment Done")
    }
    fmt.Println("Cashier getting mony from patient")
}

func (c *cashier) setNext(next department) {
    c.next = next
}

// patient.go
type patient struct {
    name string
    registrationDone bool
    doctorCheckUpDone bool
    medicineDone bool
    paymentDone bool
}

// main.go
func main() {
    cashier := &cashier{}
    
    // Set next for medical department
    medical := &medical{}
    medical.setNext(cashier)

    // Set nex for doctor  department
    doctor := &doctor{}
    doctor.setNext(medical)
    
    // Set next for reception department
    reception := &reception{}
    reception.setNext(doctor)
    patient := &patient{name: "abc"}

    // Patient visiting
    reception.execute(patient)
}
