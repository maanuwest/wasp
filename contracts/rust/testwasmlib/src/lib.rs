// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
//////// DO NOT CHANGE THIS FILE! ////////
// Change the json schema instead

use consts::*;
use testwasmlib::*;
use wasmlib::*;

mod consts;
mod testwasmlib;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_PARAM_TYPES, func_param_types_thunk);
}

//@formatter:off
pub struct FuncParamTypesParams {
    pub address:    ScImmutableAddress,
    pub agent_id:   ScImmutableAgentId,
    pub bytes:      ScImmutableBytes,
    pub chain_id:   ScImmutableChainId,
    pub color:      ScImmutableColor,
    pub hash:       ScImmutableHash,
    pub hname:      ScImmutableHname,
    pub int64:      ScImmutableInt64,
    pub request_id: ScImmutableRequestId,
    pub string:     ScImmutableString,
}
//@formatter:on

fn func_param_types_thunk(ctx: &ScFuncContext) {
    ctx.log("testwasmlib.funcParamTypes");
    let p = ctx.params();
    let params = FuncParamTypesParams {
        address: p.get_address(PARAM_ADDRESS),
        agent_id: p.get_agent_id(PARAM_AGENT_ID),
        bytes: p.get_bytes(PARAM_BYTES),
        chain_id: p.get_chain_id(PARAM_CHAIN_ID),
        color: p.get_color(PARAM_COLOR),
        hash: p.get_hash(PARAM_HASH),
        hname: p.get_hname(PARAM_HNAME),
        int64: p.get_int64(PARAM_INT64),
        request_id: p.get_request_id(PARAM_REQUEST_ID),
        string: p.get_string(PARAM_STRING),
    };
    func_param_types(ctx, &params);
    ctx.log("testwasmlib.funcParamTypes ok");
}
