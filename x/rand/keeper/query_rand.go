package keeper

import (
	"context"
	"crypto/sha256"
        "encoding/binary"
        "io"
	"strconv"
        "math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"rand/x/rand/types"
)

func (k Keeper) Rand(goCtx context.Context, req *types.QueryRandRequest) (*types.QueryRandResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
        h := sha256.New()
        blockTime := ctx.BlockHeader().Time
	unixTimestampMillis := blockTime.UnixNano()
	strTime := strconv.FormatInt(unixTimestampMillis, 10)
        validatorsHash := string(ctx.ValidatorsHash())
        appHash := string(ctx.AppHash())
        io.WriteString(h, string(strTime + validatorsHash + appHash))
        var seed uint64 = binary.BigEndian.Uint64(h.Sum(nil))
        rand.Seed(int64(seed))

        return &types.QueryRandResponse{Random: rand.Int63()}, nil
}
