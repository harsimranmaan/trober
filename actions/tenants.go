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
// Model: Singular (Tenant)
// DB Table: Plural (tenants)
// Resource: Plural (Tenants)
// Path: Plural (/tenants)
// View Template Folder: Plural (/templates/tenants/)

// TenantsResource is the resource for the Tenant model
type TenantsResource struct{
  buffalo.Resource
}

// List gets all Tenants. This function is mapped to the path
// GET /tenants
func (v TenantsResource) List(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  tenants := &models.Tenants{}

  // Paginate results. Params "page" and "per_page" control pagination.
  // Default values are "page=1" and "per_page=20".
  q := tx.PaginateFromParams(c.Params())

  // Retrieve all Tenants from the DB
  if err := q.All(tenants); err != nil {
    return err
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // Add the paginator to the context so it can be used in the template.
    c.Set("pagination", q.Paginator)

    c.Set("tenants", tenants)
    return c.Render(http.StatusOK, r.HTML("/tenants/index.plush.html"))
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(200, r.JSON(tenants))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(200, r.XML(tenants))
  }).Respond(c)
}

// Show gets the data for one Tenant. This function is mapped to
// the path GET /tenants/{tenant_id}
func (v TenantsResource) Show(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Tenant
  tenant := &models.Tenant{}

  // To find the Tenant the parameter tenant_id is used.
  if err := tx.Find(tenant, c.Param("tenant_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    c.Set("tenant", tenant)

    return c.Render(http.StatusOK, r.HTML("/tenants/show.plush.html"))
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(200, r.JSON(tenant))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(200, r.XML(tenant))
  }).Respond(c)
}

// Create adds a Tenant to the DB. This function is mapped to the
// path POST /tenants
func (v TenantsResource) Create(c buffalo.Context) error {
  // Allocate an empty Tenant
  tenant := &models.Tenant{}

  // Bind tenant to the html form elements
  if err := c.Bind(tenant); err != nil {
    return err
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(tenant)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    return responder.Wants("html", func (c buffalo.Context) error {
      // Make the errors available inside the html template
      c.Set("errors", verrs)

      // Render again the new.html template that the user can
      // correct the input.
      c.Set("tenant", tenant)

      return c.Render(http.StatusUnprocessableEntity, r.HTML("/tenants/new.plush.html"))
    }).Wants("json", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
    }).Wants("xml", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
    }).Respond(c)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a success message
    c.Flash().Add("success", T.Translate(c, "tenant.created.success"))

    // and redirect to the show page
    return c.Redirect(http.StatusSeeOther, "/tenants/%v", tenant.ID)
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusCreated, r.JSON(tenant))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusCreated, r.XML(tenant))
  }).Respond(c)
}

// Update changes a Tenant in the DB. This function is mapped to
// the path PUT /tenants/{tenant_id}
func (v TenantsResource) Update(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Tenant
  tenant := &models.Tenant{}

  if err := tx.Find(tenant, c.Param("tenant_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  // Bind Tenant to the html form elements
  if err := c.Bind(tenant); err != nil {
    return err
  }

  verrs, err := tx.ValidateAndUpdate(tenant)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    return responder.Wants("html", func (c buffalo.Context) error {
      // Make the errors available inside the html template
      c.Set("errors", verrs)

      // Render again the edit.html template that the user can
      // correct the input.
      c.Set("tenant", tenant)

      return c.Render(http.StatusUnprocessableEntity, r.HTML("/tenants/edit.plush.html"))
    }).Wants("json", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
    }).Wants("xml", func (c buffalo.Context) error {
      return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
    }).Respond(c)
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a success message
    c.Flash().Add("success", T.Translate(c, "tenant.updated.success"))

    // and redirect to the show page
    return c.Redirect(http.StatusSeeOther, "/tenants/%v", tenant.ID)
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.JSON(tenant))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.XML(tenant))
  }).Respond(c)
}

// Destroy deletes a Tenant from the DB. This function is mapped
// to the path DELETE /tenants/{tenant_id}
func (v TenantsResource) Destroy(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Tenant
  tenant := &models.Tenant{}

  // To find the Tenant the parameter tenant_id is used.
  if err := tx.Find(tenant, c.Param("tenant_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  if err := tx.Destroy(tenant); err != nil {
    return err
  }

  return responder.Wants("html", func (c buffalo.Context) error {
    // If there are no errors set a flash message
    c.Flash().Add("success", T.Translate(c, "tenant.destroyed.success"))

    // Redirect to the index page
    return c.Redirect(http.StatusSeeOther, "/tenants")
  }).Wants("json", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.JSON(tenant))
  }).Wants("xml", func (c buffalo.Context) error {
    return c.Render(http.StatusOK, r.XML(tenant))
  }).Respond(c)
}
