// This file is auto-generated, don't edit it. Thanks.

using System;
using System.Collections.Generic;
using System.IO;

using Tea;

namespace Aliyun.SDK.CCP.OSSClient.Models
{
    /**
     * search file response
     */
    public class CCPSearchFileResponse : TeaModel {
        [NameInMap("items")]
        [Validation(Required=false)]
        public List<BaseCCPFileResponse> Items { get; set; }

        [NameInMap("next_marker")]
        [Validation(Required=false)]
        public string NextMarker { get; set; }

    }

}
