// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.ccp.ccpclient.models;

import com.aliyun.tea.*;

/**
 * 
 */
public class MobileSendSmsCodeResponse extends TeaModel {
    @NameInMap("sms_code_id")
    @Validation(required = true)
    public String smsCodeId;

    public static MobileSendSmsCodeResponse build(java.util.Map<String, ?> map) throws Exception {
        MobileSendSmsCodeResponse self = new MobileSendSmsCodeResponse();
        return TeaModel.build(map, self);
    }

}
