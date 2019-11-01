// This file is auto-generated, don't edit it. Thanks.
package com.aliyun.ccp.ccpclient.models;

import com.aliyun.tea.*;

public class CCPGetUploadUrlResponse extends TeaModel {
    @NameInMap("requestId")
    @Validation(pattern = "", maxLength = 0)
    public String requestId;

    @NameInMap("create_at")
    @Validation(pattern = "", maxLength = 0)
    public String createAt;

    @NameInMap("domain_id")
    @Validation(pattern = "[a-z0-9A-Z]+", maxLength = 0)
    public String domainId;

    @NameInMap("drive_id")
    @Validation(pattern = "[0-9]+", maxLength = 0)
    public String driveId;

    @NameInMap("file_id")
    @Validation(pattern = "[a-z0-9]{1, 50}", maxLength = 0)
    public String fileId;

    @NameInMap("part_info_list")
    @Validation(pattern = "", maxLength = 0)
    public UploadPartInfo[] partInfoList;

    @NameInMap("upload_id")
    @Validation(pattern = "", maxLength = 0)
    public String uploadId;

}
