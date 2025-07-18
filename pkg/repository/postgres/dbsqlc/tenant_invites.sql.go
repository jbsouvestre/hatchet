// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: tenant_invites.sql

package dbsqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countActiveInvites = `-- name: CountActiveInvites :one
SELECT
    COUNT(*)
FROM
    "TenantInviteLink"
WHERE
    "tenantId" = $1::uuid
    AND "status" = 'PENDING'
    AND "expires" > now()
`

func (q *Queries) CountActiveInvites(ctx context.Context, db DBTX, tenantid pgtype.UUID) (int64, error) {
	row := db.QueryRow(ctx, countActiveInvites, tenantid)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createTenantInvite = `-- name: CreateTenantInvite :one
INSERT INTO "TenantInviteLink" (
    "id",
    "tenantId",
    "inviterEmail",
    "inviteeEmail",
    "expires",
    "role"
) VALUES (
    gen_random_uuid(),
    $1::uuid,
    $2::text,
    $3::text,
    $4::timestamp,
    $5::"TenantMemberRole"
) RETURNING id, "createdAt", "updatedAt", "tenantId", "inviterEmail", "inviteeEmail", expires, status, role
`

type CreateTenantInviteParams struct {
	Tenantid     pgtype.UUID      `json:"tenantid"`
	Inviteremail string           `json:"inviteremail"`
	Inviteeemail string           `json:"inviteeemail"`
	Expires      pgtype.Timestamp `json:"expires"`
	Role         TenantMemberRole `json:"role"`
}

func (q *Queries) CreateTenantInvite(ctx context.Context, db DBTX, arg CreateTenantInviteParams) (*TenantInviteLink, error) {
	row := db.QueryRow(ctx, createTenantInvite,
		arg.Tenantid,
		arg.Inviteremail,
		arg.Inviteeemail,
		arg.Expires,
		arg.Role,
	)
	var i TenantInviteLink
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TenantId,
		&i.InviterEmail,
		&i.InviteeEmail,
		&i.Expires,
		&i.Status,
		&i.Role,
	)
	return &i, err
}

const deleteTenantInvite = `-- name: DeleteTenantInvite :exec
DELETE FROM
    "TenantInviteLink"
WHERE
    "id" = $1::uuid
`

func (q *Queries) DeleteTenantInvite(ctx context.Context, db DBTX, id pgtype.UUID) error {
	_, err := db.Exec(ctx, deleteTenantInvite, id)
	return err
}

const getExistingInvite = `-- name: GetExistingInvite :one
SELECT
    id, "createdAt", "updatedAt", "tenantId", "inviterEmail", "inviteeEmail", expires, status, role
FROM
    "TenantInviteLink"
WHERE
    "inviteeEmail" = $1::text
    AND "tenantId" = $2::uuid
    AND "status" = 'PENDING'
    AND "expires" > now()
`

type GetExistingInviteParams struct {
	Inviteeemail string      `json:"inviteeemail"`
	Tenantid     pgtype.UUID `json:"tenantid"`
}

func (q *Queries) GetExistingInvite(ctx context.Context, db DBTX, arg GetExistingInviteParams) (*TenantInviteLink, error) {
	row := db.QueryRow(ctx, getExistingInvite, arg.Inviteeemail, arg.Tenantid)
	var i TenantInviteLink
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TenantId,
		&i.InviterEmail,
		&i.InviteeEmail,
		&i.Expires,
		&i.Status,
		&i.Role,
	)
	return &i, err
}

const getInviteById = `-- name: GetInviteById :one
SELECT
    id, "createdAt", "updatedAt", "tenantId", "inviterEmail", "inviteeEmail", expires, status, role
FROM
    "TenantInviteLink"
WHERE
    "id" = $1::uuid
`

func (q *Queries) GetInviteById(ctx context.Context, db DBTX, id pgtype.UUID) (*TenantInviteLink, error) {
	row := db.QueryRow(ctx, getInviteById, id)
	var i TenantInviteLink
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TenantId,
		&i.InviterEmail,
		&i.InviteeEmail,
		&i.Expires,
		&i.Status,
		&i.Role,
	)
	return &i, err
}

const listInvitesByTenantId = `-- name: ListInvitesByTenantId :many
SELECT
    id, "createdAt", "updatedAt", "tenantId", "inviterEmail", "inviteeEmail", expires, status, role
FROM
    "TenantInviteLink"
WHERE
    "tenantId" = $1::uuid
    AND (
        $2::"InviteLinkStatus" IS NULL
        OR "status" = $2::"InviteLinkStatus"
    )
    AND (
        CASE WHEN $3::boolean IS NULL THEN TRUE
        -- otherwise, if expired is true, return only expired invites, and vice versa
        ELSE $3::boolean = ("expires" < now())
        END
    )
`

type ListInvitesByTenantIdParams struct {
	Tenantid pgtype.UUID          `json:"tenantid"`
	Status   NullInviteLinkStatus `json:"status"`
	Expired  pgtype.Bool          `json:"expired"`
}

func (q *Queries) ListInvitesByTenantId(ctx context.Context, db DBTX, arg ListInvitesByTenantIdParams) ([]*TenantInviteLink, error) {
	rows, err := db.Query(ctx, listInvitesByTenantId, arg.Tenantid, arg.Status, arg.Expired)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*TenantInviteLink
	for rows.Next() {
		var i TenantInviteLink
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.TenantId,
			&i.InviterEmail,
			&i.InviteeEmail,
			&i.Expires,
			&i.Status,
			&i.Role,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTenantInvitesByEmail = `-- name: ListTenantInvitesByEmail :many
SELECT
    iv.id, iv."createdAt", iv."updatedAt", iv."tenantId", iv."inviterEmail", iv."inviteeEmail", iv.expires, iv.status, iv.role,
    t."name" as "tenantName"
FROM
    "TenantInviteLink" iv
JOIN
    "Tenant" t ON iv."tenantId" = t."id"
WHERE
    iv."inviteeEmail" = $1::text
    AND iv."status" = 'PENDING'
    AND iv."expires" > now()
`

type ListTenantInvitesByEmailRow struct {
	ID           pgtype.UUID      `json:"id"`
	CreatedAt    pgtype.Timestamp `json:"createdAt"`
	UpdatedAt    pgtype.Timestamp `json:"updatedAt"`
	TenantId     pgtype.UUID      `json:"tenantId"`
	InviterEmail string           `json:"inviterEmail"`
	InviteeEmail string           `json:"inviteeEmail"`
	Expires      pgtype.Timestamp `json:"expires"`
	Status       InviteLinkStatus `json:"status"`
	Role         TenantMemberRole `json:"role"`
	TenantName   string           `json:"tenantName"`
}

func (q *Queries) ListTenantInvitesByEmail(ctx context.Context, db DBTX, inviteeemail string) ([]*ListTenantInvitesByEmailRow, error) {
	rows, err := db.Query(ctx, listTenantInvitesByEmail, inviteeemail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*ListTenantInvitesByEmailRow
	for rows.Next() {
		var i ListTenantInvitesByEmailRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.TenantId,
			&i.InviterEmail,
			&i.InviteeEmail,
			&i.Expires,
			&i.Status,
			&i.Role,
			&i.TenantName,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTenantInvite = `-- name: UpdateTenantInvite :one
UPDATE
    "TenantInviteLink"
SET
    "status" = COALESCE($1::"InviteLinkStatus", "status"),
    "role" = COALESCE($2::"TenantMemberRole", "role")
WHERE
    "id" = $3::uuid
RETURNING id, "createdAt", "updatedAt", "tenantId", "inviterEmail", "inviteeEmail", expires, status, role
`

type UpdateTenantInviteParams struct {
	Status NullInviteLinkStatus `json:"status"`
	Role   NullTenantMemberRole `json:"role"`
	ID     pgtype.UUID          `json:"id"`
}

func (q *Queries) UpdateTenantInvite(ctx context.Context, db DBTX, arg UpdateTenantInviteParams) (*TenantInviteLink, error) {
	row := db.QueryRow(ctx, updateTenantInvite, arg.Status, arg.Role, arg.ID)
	var i TenantInviteLink
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TenantId,
		&i.InviterEmail,
		&i.InviteeEmail,
		&i.Expires,
		&i.Status,
		&i.Role,
	)
	return &i, err
}
