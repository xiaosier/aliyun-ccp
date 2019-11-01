// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.ccp.ossclient.models;

import com.aliyun.tea.*;

public class BaseCreateFileRequest extends TeaModel {
    @NameInMap("content_md5")
    @Validation(pattern = "", maxLength = 0)
    public String contentMd5;

    @NameInMap("content_type")
    @Validation(pattern = "", maxLength = 0)
    public String contentType;

    @NameInMap("name")
    @Validation(pattern = "[a-z0-9.-_]{1,1024}", maxLength = 0)
    public String name;

    @NameInMap("part_info_list")
    @Validation(pattern = "", maxLength = 0)
    public UploadPartInfo[] partInfoList;

    @NameInMap("size")
    @Validation(pattern = "", maxLength = 0)
    public Long size;

    @NameInMap("type")
    @Validation(pattern = "", maxLength = 0)
    public String type;

}
