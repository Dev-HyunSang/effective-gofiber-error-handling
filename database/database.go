package database

import (
	"context"
	"errors"
	"github.com/dev-hyunsang/effective-gofiber-error-handling/cmd"
	"github.com/dev-hyunsang/effective-gofiber-error-handling/ent"
	"github.com/dev-hyunsang/effective-gofiber-error-handling/ent/todo"
	"github.com/google/uuid"
	"time"
)

var (
	ErrNotFoundToDo = errors.New("존재하지 않는 할일 항목입니다.")

	// Relevant DataBase Error
	ErrFailedConnectDataBase = errors.New("정상적으로 데이터베이스에 접속하지 못 했습니다. 확인 후 다시 시도하세요.")
	ErrFailedCreateDataBase  = errors.New("정상적으로 스키마 리소스를 만들지 못 했습니다.")
	ErrFailedInsertDataBase  = errors.New("정상적으로 삽입되지 않았습니다. 다시 시도해 주세요.")
	ErrFailedUpdateDataBase  = errors.New("정상적으로 수정되지 않았습니다. 다시 시도해 주세요.")
	ErrFailedReadDataBase    = errors.New("정상적으로 읽지 못 했어요. 다시 시도해 주세요.")
	ErrFailedDeleteDataBase  = errors.New("정상적으로 삭제하지 못 했어요. 다시 시도해 주세요.")
)

func ConnectionSQLite() (*ent.Client, error) {
	client, err := ent.Open("sqlite3", "file:todo.db?_fk=1")
	if err != nil {
		return nil, err
	}

	defer client.Close()

	if err = client.Schema.Create(context.Background()); err != nil {
		return nil, ErrFailedCreateDataBase
	}

	return client, nil
}

// 새로운 ToDo를 만듭니다.
// cmd.ToDo 구조체를 입력으로 필요로 하고, ent.ToDo 구조체와 오류를 반환합니다.
func Create(todo cmd.ToDo) (*ent.ToDo, error) {
	client, err := ConnectionSQLite()
	if err != nil {
		return nil, ErrFailedConnectDataBase
	}

	u, err := client.ToDo.Create().
		SetTodoUUID(todo.ToDoUUID).
		SetTodoContext(todo.ToDoContext).
		SetUpdatedAt(todo.UpdatedAt).
		SetCreatedAt(todo.CreatedAt).
		Save(context.Background())

	if err != nil {
		return nil, ErrFailedInsertDataBase
	}

	return u, nil
}

// 생성되어 있는 모든 항목들에 대해서 조회하는 함수
// 입력은 없으며, ToDo 구조체와 오류를 반환함.
func AllToDoRead() ([]*ent.ToDo, error) {
	client, err := ConnectionSQLite()
	if err != nil {
		return nil, ErrFailedConnectDataBase
	}

	date, err := client.ToDo.Query().
		All(context.Background())
	if err != nil {
		return nil, ErrFailedReadDataBase
	}

	return date, nil
}

func ParticularToDoRead(todoUUID uuid.UUID) ([]*ent.ToDo, error) {
	client, err := ConnectionSQLite()
	if err != nil {
		return nil, ErrFailedConnectDataBase
	}

	data, err := client.ToDo.Query().
		Where(todo.TodoUUID(todoUUID)).
		All(context.Background())
	if err != nil {
		return nil, ErrFailedReadDataBase
	}

	return data, nil
}

// 기존의 생성되어 있는 ToDo를 수정합니다.
// cmd.ToDo 객체가 필요하며, 결과값인 정수와 오류를 반환합니다.
func Update(data *cmd.ToDo) (int, error) {
	client, err := ConnectionSQLite()
	if err != nil {
		return 0, ErrFailedConnectDataBase
	}

	u, err := client.ToDo.
		Update().
		Where(todo.TodoUUID(data.ToDoUUID)).
		SetTodoContext(data.ToDoContext).
		SetUpdatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		return 0, ErrFailedUpdateDataBase
	}

	return u, nil
}

// 생성되어 있는 ToDo를 삭제하는 용도로 사용합니다.
// 생성되어 있는 ToDo의 UUID를 필요로 하며, 결과값인 정수와 오류를 반환합니다.
func Delete(todoUUID uuid.UUID) (int, error) {
	client, err := ConnectionSQLite()
	if err != nil {
		return 0, ErrFailedConnectDataBase
	}

	u := client.ToDo.
		Delete().
		Where(todo.TodoUUID(todoUUID)).
		ExecX(context.Background())

	return u, nil
}
