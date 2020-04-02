package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/gomodule/redigo/redis"

	"github.com/cjcjcj/todo/todo/domains"

	"github.com/rafaeljusto/redigomock"
)

func TestRedisTodoCreate(t *testing.T) {
	var (
		id    uint = 1
		title      = "test"
	)

	rc := redigomock.NewConn()
	h := NewRedisTodoRepository(rc)

	t.Run("OK", func(t *testing.T) {
		cmdInc := rc.Command("INCR", redisIDfield).Expect(int64(id))
		cmdAdd := rc.Command("HSET", redisTODOsField, id, fmt.Sprintf(`{"id":%d,"title":"%s","closed":false}`, id, title)).Expect("OK")

		te := domains.NewTodoFromString(title)

		err := h.Create(context.TODO(), te)
		if err != nil {
			t.Fatal(err)
		}

		if rc.Stats(cmdInc) != 1 {
			t.Fatal("INC command not called")
		}
		if rc.Stats(cmdAdd) != 1 {
			t.Fatal("HSET command not called")
		}

		if te.ID != id {
			t.Errorf("te.ID != expected id(%d != %d)", te.ID, id)
		}
		if te.Title != title {
			t.Errorf("te.Title != expected title(%s != %s)", te.Title, title)
		}
		if te.Closed != false {
			t.Errorf("te.Closed != expected closed(%v != %v)", te.Closed, false)
		}
	})

	t.Run("error-incr", func(t *testing.T) {
		rc.Command("INCR", redisIDfield).ExpectError(fmt.Errorf("HSET error"))

		te := domains.NewTodoFromString(title)

		err := h.Create(context.TODO(), te)
		if err == nil {
			t.Fatal("Should return an error")
		}
	})

	t.Run("error-hset", func(t *testing.T) {
		rc.Command("INCR", redisIDfield).Expect(int64(id))
		rc.Command("HSET", redisTODOsField, id, fmt.Sprintf(`{"id":%d,"title":"%s","closed":false}`, id, title)).ExpectError(fmt.Errorf("HSET error"))

		te := domains.NewTodoFromString(title)

		err := h.Create(context.TODO(), te)
		if err == nil {
			t.Fatal("Should return an error")
		}
	})
}

func TestRedisTodoUpdate(t *testing.T) {
	var (
		id     uint = 1
		title       = "test"
		closed      = true
	)

	rc := redigomock.NewConn()
	h := NewRedisTodoRepository(rc)

	t.Run("OK", func(t *testing.T) {
		cmd := rc.Command("HSET", redisTODOsField, id, fmt.Sprintf(`{"id":%d,"title":"%s","closed":%v}`, id, title, closed)).Expect("OK")

		te := &domains.Todo{ID: id, Title: title, Closed: closed}

		if err := h.Update(context.TODO(), te); err != nil {
			t.Fatal(err)
		}

		if rc.Stats(cmd) != 1 {
			t.Fatal("HSET not executed")
		}

		if te.ID != id {
			t.Errorf("te.ID != expected id(%d != %d)", te.ID, id)
		}
		if te.Title != title {
			t.Errorf("te.Title != expected title(%s != %s)", te.Title, title)
		}
		if te.Closed != closed {
			t.Errorf("te.Closed != expected closed(%v != %v)", te.Closed, closed)
		}
	})

	t.Run("error", func(t *testing.T) {
		rc.Command("HSET", redisTODOsField, id, fmt.Sprintf(`{"id":%d,"title":"%s","closed":%v}`, id, title, closed)).ExpectError(fmt.Errorf("expected error"))

		te := &domains.Todo{ID: id, Title: title, Closed: closed}

		if err := h.Update(context.TODO(), te); err == nil {
			t.Fatal(err)
		}
	})
}

func TestRedisTodoDelete(t *testing.T) {
	var id uint = 1

	rc := redigomock.NewConn()
	h := NewRedisTodoRepository(rc)

	t.Run("OK", func(t *testing.T) {
		// HDEL returns count of removed fields
		cmd := rc.Command("HDEL", redisTODOsField, id).Expect(1)

		if err := h.Delete(context.TODO(), id); err != nil {
			t.Fatal(err)
		}

		if rc.Stats(cmd) != 1 {
			t.Fatal("HDEL not executed")
		}
	})

	t.Run("error", func(t *testing.T) {
		rc.Command("HDEL", redisTODOsField, id).ExpectError(fmt.Errorf("expected error"))

		if err := h.Delete(context.TODO(), id); err == nil {
			t.Fatal("Should return an error")
		}
	})
}

func TestRedisTodoGetByID(t *testing.T) {
	var (
		id    uint = 1
		title      = "test"
	)

	rc := redigomock.NewConn()
	h := NewRedisTodoRepository(rc)

	t.Run("OK", func(t *testing.T) {
		cmdGet := rc.Command("HGET", redisTODOsField, id).Expect(fmt.Sprintf(`{"id":%d,"title":"%s","closed":false}`, id, title))

		te, err := h.GetByID(context.TODO(), id)
		if err != nil {
			t.Fatal(err)
		}

		if rc.Stats(cmdGet) != 1 {
			t.Fatal("HGET not executed")
		}

		if te.ID != 1 {
			t.Errorf("te.ID != expected id(%d != %d)", te.ID, id)
		}
		if te.Title != title {
			t.Errorf("te.Title != expected title(%s != %s)", te.Title, title)
		}
		if te.Closed != false {
			t.Errorf("te.Closed != expected closed(%v != %v)", te.Closed, false)
		}
	})

	t.Run("not-found", func(t *testing.T) {
		rc.Command("HGET", redisTODOsField, id).ExpectError(redis.ErrNil)

		v, err := h.GetByID(context.TODO(), id)
		if err != nil {
			t.Fatal(err)
		}
		if v != nil {
			t.Fatalf("%v should be nil", v)
		}
	})

	t.Run("error", func(t *testing.T) {
		rc.Command("HGET", redisTODOsField, id).ExpectError(fmt.Errorf("expected error"))

		_, err := h.GetByID(context.TODO(), id)
		if err == nil {
			t.Fatal("Should return an error")
		}
	})
}

func TestRedisTodoGetAll(t *testing.T) {
	var (
		id    uint = 1
		title      = "test"
	)

	rc := redigomock.NewConn()
	h := NewRedisTodoRepository(rc)

	t.Run("OK", func(t *testing.T) {
		cmdGet := rc.Command("HVALS", redisTODOsField).ExpectStringSlice(
			fmt.Sprintf(`{"id":%d,"title":"%s","closed":false}`, id, title),
		)

		te, err := h.GetAll(context.TODO())
		if err != nil {
			t.Fatal(err)
		}

		if rc.Stats(cmdGet) != 1 {
			t.Fatal("HVALS not executed")
		}

		if te[0].ID != 1 {
			t.Errorf("te[0].ID != expected id(%d != %d)", te[0].ID, id)
		}
		if te[0].Title != title {
			t.Errorf("te[0].Title != expected title(%s != %s)", te[0].Title, title)
		}
		if te[0].Closed != false {
			t.Errorf("te[0].Closed != expected closed(%v != %v)", te[0].Closed, false)
		}
	})

	t.Run("error", func(t *testing.T) {
		rc.Command("HVALS", redisTODOsField).ExpectError(fmt.Errorf("expected error"))

		_, err := h.GetAll(context.TODO())
		if err == nil {
			t.Fatal("Should return an error")
		}
	})
}
