package actions

import (

  "fmt"
  "net/http"
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/pop/v5"
  "github.com/gobuffalo/x/responder"
  "trober/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Carrier)
// DB Table: Plural (carriers)
// Resource: Plural (Carriers)
// Path: Plural (/carriers)
// View Template Folder: Plural (/templates/carriers/)

// CarriersResource is the resource for the Carrier model
type CarriersResource struct{
  buffalo.Resource
}

// List gets all Carriers. This function is mapped to the path
// GET /carriers
func (v CarriersResource) List(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  carriers := &models.Carriers{}

  // Paginate results. Params "page" and "per_page" control pagination.
  // Default values are "page=1" and "per_page=20".
  q := tx.PaginateFromParams(c.Params())

  // Retrieve all Carriers from the DB
  if err := q.All(carriers); err != nil {
    return err
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // Add the paginator to the context so it can be used in the template.
    c.Set("pagination", q.Paginator)

    c.Set("carriers", carriers)
    return c.Render(http.StatusOK, r.HTML("/carriers/index.plush.html"))
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(200, r.JSON(carriers))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(200, r.XML(carriers))
  }).Respond(c)
}

// Show gets the data for one Carrier. This function is mapped to
// the path GET /carriers/{carrier_id}
func (v CarriersResource) Show(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Carrier
  carrier := &models.Carrier{}

  // To find the Carrier the parameter carrier_id is used.
  if err := tx.Find(carrier, c.Param("carrier_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    c.Set("carrier", carrier)

    return c.Render(http.StatusOK, r.HTML("/carriers/show.plush.html"))
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(200, r.JSON(carrier))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(200, r.XML(carrier))
  }).Respond(c)
}

// Create adds a Carrier to the DB. This function is mapped to the
// path POST /carriers
func (v CarriersResource) Create(c buffalo.Context) error {
  // Allocate an empty Carrier
  carrier := &models.Carrier{}

  // Bind carrier to the html form elements
  if err := c.Bind(carrier); err != nil {
    return err
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(carrier)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    return responder.Wants("html", func (c buffalo.Context) error {
      // Make the errors available inside the html template
      c.Set("errors", verrs)

      // Render again the new.html template that the user can
      // correct the input.
      c.Set("carrier", carrier)

      return c.Render(http.StatusUnprocessableEntity, r.HTML("/carriers/new.plush.html"))
    }).Wants("json", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
    }).Wants("xml", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
    }).Respond(c)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a success message
    c.Flash().Add("success", T.Translate(c, "carrier.created.success"))

    // and redirect to the show page
    return c.Redirect(http.StatusSeeOther, "/carriers/%v", carrier.ID)
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusCreated, r.JSON(carrier))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusCreated, r.XML(carrier))
  }).Respond(c)
}

// Update changes a Carrier in the DB. This function is mapped to
// the path PUT /carriers/{carrier_id}
func (v CarriersResource) Update(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Carrier
  carrier := &models.Carrier{}

  if err := tx.Find(carrier, c.Param("carrier_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  // Bind Carrier to the html form elements
  if err := c.Bind(carrier); err != nil {
    return err
  }

  verrs, err := tx.ValidateAndUpdate(carrier)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    return responder.Wants("html", func (c buffalo.Context) error {
      // Make the errors available inside the html template
      c.Set("errors", verrs)

      // Render again the edit.html template that the user can
      // correct the input.
      c.Set("carrier", carrier)

      return c.Render(http.StatusUnprocessableEntity, r.HTML("/carriers/edit.plush.html"))
    }).Wants("json", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
    }).Wants("xml", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
    }).Respond(c)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a success message
    c.Flash().Add("success", T.Translate(c, "carrier.updated.success"))

    // and redirect to the show page
    return c.Redirect(http.StatusSeeOther, "/carriers/%v", carrier.ID)
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.JSON(carrier))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.XML(carrier))
  }).Respond(c)
}

// Destroy deletes a Carrier from the DB. This function is mapped
// to the path DELETE /carriers/{carrier_id}
func (v CarriersResource) Destroy(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Carrier
  carrier := &models.Carrier{}

  // To find the Carrier the parameter carrier_id is used.
  if err := tx.Find(carrier, c.Param("carrier_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  if err := tx.Destroy(carrier); err != nil {
    return err
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a flash message
    c.Flash().Add("success", T.Translate(c, "carrier.destroyed.success"))

    // Redirect to the index page
    return c.Redirect(http.StatusSeeOther, "/carriers")
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.JSON(carrier))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.XML(carrier))
  }).Respond(c)
}