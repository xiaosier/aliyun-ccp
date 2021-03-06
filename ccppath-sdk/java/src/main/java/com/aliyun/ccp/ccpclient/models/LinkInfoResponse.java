// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.ccp.ccpclient.models;

import com.aliyun.tea.*;

/**
 * 
 */
public class LinkInfoResponse extends TeaModel {
    @NameInMap("authentication_type")
    @Validation(required = true)
    public String authenticationType;

    @NameInMap("created_at")
    @Validation(required = true)
    public Long createdAt;

    @NameInMap("domain_id")
    @Validation(required = true)
    public String domainId;

    @NameInMap("identity")
    @Validation(required = true)
    public String identity;

    @NameInMap("last_login_time")
    @Validation(required = true)
    public Long lastLoginTime;

    @NameInMap("status")
    @Validation(required = true)
    public String status;

    @NameInMap("user_id")
    @Validation(required = true)
    public String userId;

    public static LinkInfoResponse build(java.util.Map<String, ?> map) throws Exception {
        LinkInfoResponse self = new LinkInfoResponse();
        return TeaModel.build(map, self);
    }

}
